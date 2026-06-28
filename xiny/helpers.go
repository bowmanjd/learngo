package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

// logErr prints a contextual message with the error if err is non-nil.
// Returns true when an error was present, so callers can short-circuit.
func logErr(msg string, err error) bool {
	if err != nil {
		fmt.Println(msg, err)
		return true
	}
	return false
}

// writeResponse writes a formatted string to an http.ResponseWriter,
// logging any write failure via structured logging.
func writeResponse(w http.ResponseWriter, format string, args ...any) {
	if _, err := fmt.Fprintf(w, format, args...); err != nil {
		slog.Default().Error("failed to write response", "err", err)
	}
}

// startServer launches an HTTP server in a background goroutine and returns
// a channel that receives any startup error (excluding ErrServerClosed).
func startServer(srv *http.Server, logger *slog.Logger) <-chan error {
	errCh := make(chan error, 1)
	go func() {
		logger.Info("starting demo server", "addr", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
		close(errCh)
	}()
	return errCh
}

// shutdownServer gracefully shuts down the server within the given timeout.
func shutdownServer(srv *http.Server, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return srv.Shutdown(ctx)
}
