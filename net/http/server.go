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
)

// ListenAndServeUDS listens for HTTP requests over a UNIX Domain Socket on p.
func ListenAndServeUDS(s *http.Server, p string) error {
	uds, err := net.Listen("unix", p)
	if err != nil {
		return fmt.Errorf("cannot create socket: %w", err)
	}

	return s.Serve(uds) //nolint:wrapcheck
}

// ShutdownServerOn listen for sig signal and gracefully shuts down the given
// server, ctx is a function that provides a context to be used during server
// shutdown.
func ShutdownServerOn(
	s *http.Server,
	sig os.Signal,
	ctx func() context.Context,
) {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, sig)
		<-c

		if err := s.Shutdown(ctx()); err != nil {
			s.ErrorLog.Print(err)
		}
	}()
}
