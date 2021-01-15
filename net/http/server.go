// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package http

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

// Options wraps all the customizable options for a Server.
type Options struct {
	// TCP address to listen on. If a file path is given, the server will use a
	// Unix Domain Socket.
	Addr string

	// Cert and Key are required for enabling TLS.
	Cert, Key string

	// Requests handler. Just as http.Server, if nil is given,
	// http.DefaultServeMux will be used.
	Handler http.Handler

	ShutdownCtx func() context.Context
}

// Server is a http.Server with some extra functionalities.
type Server struct {
	http.Server

	// TLS certificates.
	Cert, Key string

	// Shutdown context used for gracefully shutdown, it is implemented as a
	// function since deadlines will start at server creation and not at shutdown.
	ShutdownCtx func() context.Context

	// Gracefully shutdown done notifier.
	done chan struct{}
}

// NewServer creates a new Server.
func NewServer(opts Options) *Server {
	s := new(Server)
	s.Addr = opts.Addr
	s.Cert = opts.Cert
	s.Key = opts.Key
	s.done = make(chan struct{})
	s.Handler = opts.Handler
	s.ShutdownCtx = defaultShutdownCtx

	if opts.ShutdownCtx != nil {
		s.ShutdownCtx = opts.ShutdownCtx
	}

	return s
}

// ListenAndServe listens for HTTP requests. If s.Addr is a filepath, the
// server will listen in a Unix Domain Socket; if s.Cert and s.Key are not
// empty, the server will listen over a TLS listener; otherwise the server will
// listen on a regular TCP listener.
//
// By default, the server will gracefully shutdown when the program receives a
// SIGINT signal.
func (s *Server) ListenAndServe() (err error) {
	go s.gracefullyShutdown()

	switch {
	case !strings.Contains(s.Addr, ":"):
		uds, udsErr := net.Listen("unix", s.Addr)
		if udsErr != nil {
			return fmt.Errorf("cannot create socket: %w", udsErr)
		}

		err = s.Server.Serve(uds)
	case s.Cert != "" && s.Key != "":
		err = s.Server.ListenAndServeTLS(s.Cert, s.Key)
	default:
		err = s.Server.ListenAndServe()
	}

	<-s.done

	return fmt.Errorf("server stopped: %w", err)
}

// gracefullyShutdown starts a gracefully shutdown when a SIGTERM signal is
// launched.
func (s *Server) gracefullyShutdown() {
	defer close(s.done)

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	if err := s.Shutdown(s.ShutdownCtx()); err != nil {
		s.Server.ErrorLog.Print(err)
	}
}

func defaultShutdownCtx() context.Context { return context.Background() }
