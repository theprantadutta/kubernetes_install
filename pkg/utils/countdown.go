package utils

import (
	"fmt"
	"time"

	"github.com/theprantadutta/kubernetes_install/pkg/logger"
)

// Countdown function with customizable message
func Countdown(seconds int, message string) {
	log := logger.New()

	// Log the custom message at the start
	log.Info(message)

	// Countdown logic with logging
	for i := seconds; i > 0; i-- {
		fmt.Printf("\r%d seconds...", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Print("\r") // Clear the countdown line

	// Log after the countdown finishes
	log.Info("Press Ctrl+C key to exit anytime.")
}
