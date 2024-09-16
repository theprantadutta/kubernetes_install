package utils

import (
	"io"
	"os"
	"os/exec"

	"github.com/theprantadutta/kubernetes_install/pkg/logger"
)

// RunCommand runs a terminal command using Bash and logs the output/error in real-time
func RunCommand(command string, args ...string) {
	log := logger.New() // Create a new logger instance

	log.Info("Starting execution of command: " + command)

	// Combine the command and its arguments into a single string for Bash execution
	cmdArgs := append([]string{"-c", command}, args...)

	// Create the command to be executed in Bash
	cmd := exec.Command("bash", cmdArgs...)

	// Get a pipe to both stdout and stderr to capture output in real-time
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Error("Failed to get StdoutPipe: " + err.Error())
		os.Exit(1)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Error("Failed to get StderrPipe: " + err.Error())
		os.Exit(1)
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		log.Error("Failed to start the command: " + err.Error())
		os.Exit(1)
	}

	// Capture and display the output in real-time
	go func() {
		_, _ = io.Copy(os.Stdout, stdout) // Copy stdout to terminal
	}()
	go func() {
		_, _ = io.Copy(os.Stderr, stderr) // Copy stderr to terminal
	}()

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		log.Error("Command failed: " + err.Error())
		os.Exit(1)
	}

	log.Success("Finished executing command: " + command)
}
