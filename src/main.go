package main

import (
	"fmt"
	"os"

	"github.com/theprantadutta/kubernetes_install/pkg/installer"
	"github.com/theprantadutta/kubernetes_install/pkg/logger"
	"github.com/theprantadutta/kubernetes_install/pkg/utils"
)

// MoveCursorUp moves the cursor up by the specified number of lines
func MoveCursorUp(lines int) {
	fmt.Printf("\033[%dA", lines)
}

// UpdateProgressLine prints or updates a line at a specific position
func UpdateProgressLine(lineNumber int, step, total int, message string) {
	log := logger.New()
	MoveCursorUp(lineNumber)
	log.Info("#### Step %d of %d ####: %s\n", step, total, message)
	// Move cursor back down to the original position
	fmt.Printf("\033[%dB", lineNumber)
}

func main() {
	log := logger.New()

	// Countdown before starting the script
	utils.Countdown(5, "Starting the Kubernetes Install Script in...")

	if !utils.HasRootPrivileges() {
		log.Error("Not running with root (sudo) privileges. Exiting...")
		os.Exit(1)
	}
	log.Success("Running with root (sudo) privileges.")

	totalSteps := 6
	progressLineNumber := 2 // Set this to the line number where you want the progress to appear

	// Print initial placeholder lines to position the progress line
	for i := 0; i < totalSteps; i++ {
		fmt.Println()
	}

	// Update the progress line for the first step
	UpdateProgressLine(progressLineNumber, 1, totalSteps, "Initializing Install Script...")

	// Perform step 2
	utils.Countdown(5, "Adding Load Balancer IP in...")
	installer.AddLoadBalancerIp()
	UpdateProgressLine(progressLineNumber, 2, totalSteps, "Added Load Balancer IP")

	// Perform step 3
	utils.Countdown(5, "Making System Ready to install in...")
	installer.ReadySystem()
	UpdateProgressLine(progressLineNumber, 3, totalSteps, "System Ready")

	// Perform step 4
	utils.Countdown(5, "Making Kernel Ready for Containerd && Kubernetes in...")
	installer.ReadyKernel()
	UpdateProgressLine(progressLineNumber, 4, totalSteps, "Kernel Ready")

	// Perform step 5
	utils.Countdown(5, "Installing Containerd in...")
	installer.InstallContainerd()
	UpdateProgressLine(progressLineNumber, 5, totalSteps, "Containerd Installed")

	// Perform step 6
	utils.Countdown(5, "Installing Kubernetes in...")
	installer.InstallKubernetes()
	UpdateProgressLine(progressLineNumber, 6, totalSteps, "Kubernetes Installed")

	// Log completion
	log.Info("Finished Running the script successfully.")
}

// c := "ls -lh /f/Development/KDS/kubernetes_install"
// utils.RunCommand(c)
