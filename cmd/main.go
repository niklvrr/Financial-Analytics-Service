package main

import (
	"context"
	"github.com/niklvrr/Financial-Analytics-Service/internal/app"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	application := app.NewApp(ctx)

	setupGracefulShutdown(application)

	application.Run()
}

func setupGracefulShutdown(app *app.App) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		sig := <-sigChan
		log.Printf("Получен сигнал завершения: %s", sig.String())
		app.Stop()
		os.Exit(0)
	}()
}
