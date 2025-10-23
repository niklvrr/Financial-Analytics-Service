package main

import (
	"github.com/niklvrr/FinancialAnalyticsService/internal/app"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	application := app.NewApp()

	setupGracefulShutdown(application)

	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}

func setupGracefulShutdown(app *app.App) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		sig := <-sigChan
		log.Println("Получен сигнал завершения: %v", sig)
		app.Stop()
		os.Exit(0)
	}()
}
