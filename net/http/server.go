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
	// TCP address to listen on. If a path to a file is given, the server will use
	// a Unix Domain Socket.
	Addr string

	// Requests handler. Just as http.Server, if nil is given,
	// http.DefaultServeMux will be used.
	Handler http.Handler

	Log         *log.Logger
	ELog        *log.Logger
	ShutdownCtx func() context.Context
}

// Server is a http.Server with some extra functionalities.
type Server struct {
	http.Server

	// Logger for regular logging, if nil is given, stdout will be used.
	Log *log.Logger

	// Logger for error logging, if nil is given, stderr will be used.
	ELog *log.Logger

	// Shutdown context used for gracefully shutdown, it is implemented as a
	// function since deadlines will start at server creation and not at shutdown.
	ShutdownCtx func() context.Context

	// Gracefully shutdown done notifier.
	Done chan struct{}
}

// NewServer creates and setups a new Server.
func NewServer(c Config) *Server {
	s := new(Server)
	s.Done = make(chan struct{})
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
			s.ELog.Fatalf("Can't close the server gracefully.\n%v", err)
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
	if c.Addr != "" {
		s.Addr = c.Addr
	}

	if c.Handler != nil {
		s.Handler = c.Handler
	}

	if c.Log != nil {
		s.Log = c.Log
	} else if s.Log == nil {
		s.Log = log.New(os.Stdout, ServerLogPrefix, log.LstdFlags)
	}

	if c.ELog != nil {
		s.ELog = c.ELog
	} else if s.ELog == nil {
		s.ELog = log.New(os.Stderr, ServerLogPrefix, log.LstdFlags)
	}

	if c.ShutdownCtx != nil {
		s.ShutdownCtx = c.ShutdownCtx
	} else if s.ShutdownCtx == nil {
		s.ShutdownCtx = func() context.Context { return context.Background() }
	}
}
