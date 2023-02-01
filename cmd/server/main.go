package main

import (
	"hexa-example-go/internal/container"
	"hexa-example-go/internal/server"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	cont := container.New()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	go handleKillSignal(interrupt, cont.Logger)

	r := server.New(cont)
	r.Run(cont.Config.Server.Host + ":" + cont.Config.Server.Port)
}

func handleKillSignal(interrupt chan os.Signal, logger zap.Logger) {
	for {
		killSignal := <-interrupt
		// TODO: How do we want to handle signals? what should we close, wait for it to end?
		switch killSignal {
		case syscall.SIGINT:
			logger.Info("Got SIGINT...")
			os.Exit(0)
		case syscall.SIGTERM:
			logger.Info("Got SIGTERM...")
			os.Exit(0)
		}
	}
}
