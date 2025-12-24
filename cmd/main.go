package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"project/internal/config"
	"project/internal/logger"
	"syscall"
)

func start(ctx context.Context) {
	// Load config for app
	if err := config.Load(); err != nil {
		log.Fatalln("Error when load config: ", err)
	}
	// Initialize the logger
	logger, err := logger.Init(config.Config.App.Mode)
	if err != nil {
		log.Fatalln("Error when init logger: ", err)
	}

	logger.Info("Service successfully started...")
}

func main() {
	// Create channels and context for signals and errors
	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal, 1)
	errCh := make(chan error, 1)
	defer cancel()
	// Set up signal handling
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)
	// Start the main logic in a separate goroutine
	go func() {
		defer func() {
			if r := recover(); r != nil {
				errCh <- fmt.Errorf("panic recovered: %v", r)
			}
		}()
		// ----------------------------------------
		// ----------- START MAIN LOGIC -----------
		// ----------------------------------------
		start(ctx)
		// ----------------------------------------
		// ----------------------------------------
		// ----------------------------------------
	}()
	// Wait for shutdown signals or errors
	select {
	case sig := <-quit:
		log.Printf("Received signal: %v, initiating graceful shutdown", sig)
	case err := <-errCh:
		log.Printf("Error occurred: %v, initiating shutdown", err)
	}
	// Cancel context to signal all goroutines to stop
	cancel()
	log.Println("Graceful shutdown complete")
}
