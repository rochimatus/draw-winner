package main

import (
	"context"
	"github.com/rochimatus/draw-winner/configuration"
	"github.com/rochimatus/draw-winner/logger"
	"github.com/rochimatus/draw-winner/server"
	"github.com/rochimatus/draw-winner/server/http/handler"
	"github.com/rochimatus/draw-winner/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	gracefulShutdownCtx := handleGracefulShutdown()

	config, _ := configuration.LoadConfig("./configuration")

	servicer, err := service.New()
	panicIfError(err, "unable to initialize service")

	handler, err := handler.New(servicer)
	panicIfError(err, "unable to initialize handler")

	httpServer := http.Server{
		Addr:              ":" + config.HTTP.Port,
		ReadHeaderTimeout: config.HTTP.ReadHeaderTimeout,
	}

	// starting the http server on a separate goroutine
	go func() {
		err = server.Start(gracefulShutdownCtx, handler, &httpServer)
		panicIfError(err, "error happened while starting the server. Terminating app.")
	}()

	// gracefully shutting down all processes & releasing all resources upon receiving SIGINT/SIGTERM signal
	<-gracefulShutdownCtx.Done()

	serverShutdownContext, serverShutdownCancel := context.WithTimeout(context.Background(), config.HTTP.GraceFulShutDownDuration)
	defer serverShutdownCancel()
	server.Shutdown(serverShutdownContext, &httpServer)
}

func handleGracefulShutdown() context.Context {
	gracefulShutdownChannel := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdownChannel, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		osCall := <-gracefulShutdownChannel

		logger.Info("Received system call", "os_call", osCall.String())
		cancel()
	}()

	return ctx
}

func panicIfError(err error, message string) {
	if err != nil {
		logger.Error(err, message)
		panic(err)
	}
}
