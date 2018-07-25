// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package http

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

// ServerLogPrefix is the default server loggers prefix.
const ServerLogPrefix = "[SERVER] "

// Config wraps all the customizable options from Server.
type Config struct {
	Addr        string
	Handler     http.Handler
	Logger      *log.Logger
	ErrLogger   *log.Logger
	Ctx         context.Context
	ShutdownCtx func() context.Context
	Done        chan struct{}
}

// Server is a http.Server with extra functionalities.
type Server struct {
	http.Server

	Log         *log.Logger
	ELog        *log.Logger
	Ctx         context.Context
	ShutdownCtx func() context.Context
	Done        chan struct{}
}

// NewServer creates and setups a new Server.
func NewServer(c Config) *Server {
	s := new(Server)
	s.Setup(c)
	return s
}

// ListenAndServe starts listening in a TCP address or in a Unix Domain Socket.
func (s *Server) ListenAndServe() error {
	addr := s.Addr

	// Gracefully shutdown
	go func() {
		defer close(s.Done)

		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		s.Log.Println("Interrupt signal received, shutting down the server..")

		if err := s.Shutdown(s.ShutdownCtx()); err != nil {
			s.ELog.Printf("Can't close the server gracefully.\n%v", err)
		} else {
			s.Log.Println("All the pending tasks were done.")
			s.Log.Println("Server closed.")
		}
	}()

	if strings.Contains(addr, "/") {
		uds, err := net.Listen("unix", addr)

		if err != nil {
			s.ELog.Printf("Can't use the socket %s.\n%v", addr, err)
			return err
		}

		s.Log.Printf("Using UDS %v..\n", addr)
		return s.Server.Serve(uds)
	}

	s.Log.Printf("Listening on %v..\n", addr)
	return s.Server.ListenAndServe()
}

// Setup prepares the Server with the given Config.
func (s *Server) Setup(c Config) {
	s.Addr = c.Addr
	s.Handler = c.Handler
	s.Ctx = c.Ctx

	if c.Logger == nil {
		s.Log = log.New(os.Stdout, ServerLogPrefix, log.LstdFlags)
	} else {
		s.Log = c.Logger
	}

	if c.ErrLogger == nil {
		s.ELog = log.New(os.Stderr, ServerLogPrefix, log.LstdFlags)
	} else {
		s.ELog = c.ErrLogger
	}

	if c.ShutdownCtx == nil {
		s.ShutdownCtx = func() context.Context { return context.Background() }
	} else {
		s.ShutdownCtx = c.ShutdownCtx
	}

	if c.Done == nil {
		s.Done = make(chan struct{})
	} else {
		s.Done = c.Done
	}
}
