package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/coryt/anagram/internal/application"
	"github.com/coryt/anagram/internal/ports"
)

const (
	defaultDebugMode bool = true
)

func main() {
	ctx := context.Background()
	application := application.NewApplication(ctx, loadEnvConfig())
	httpServer := ports.ServeWeb(application)

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = httpServer.Shutdown()

	fmt.Println("Successfully shutdown")
}

func loadEnvConfig() bool {
	debugEnv := os.Getenv("debug")
	if debugEnv == "" {
		return defaultDebugMode
	}

	debugMode, err := strconv.ParseBool(debugEnv)
	if err != nil {
		return defaultDebugMode
	}

	return debugMode
}
