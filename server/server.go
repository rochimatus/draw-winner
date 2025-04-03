package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/rochimatus/draw-winner/logger"
	"log/slog"
	"net/http"
)

type Service interface {
}

// Start starts the http server based on the given configuration.
func Start(ctx context.Context, handler Handler, httpServer *http.Server) error {
	slog.Info("Starting the Server...", "port", httpServer.Addr)

	router, err := NewRouter(handler)
	if err != nil {
		logger.Error(err, "error encountered while creating routes")

		return fmt.Errorf("unable to create routes: %w", err)
	}

	httpServer.Handler = router

	httpErr := httpServer.ListenAndServe()
	if httpErr != nil && !errors.Is(httpErr, http.ErrServerClosed) {
		logger.Error(httpErr, "error starting HTTP server")

		return fmt.Errorf("error starting HTTP server: %w", httpErr)
	}

	return nil
}

// Shutdown shuts down the http server.
func Shutdown(ctx context.Context, httpServer *http.Server) {
	if serverErr := httpServer.Shutdown(ctx); serverErr != nil {
		//	log.WithError(serverErr).Error("error happened while shutting down the server.")

		return
	}

	logger.Info("Server shutdown properly!!!")
}
