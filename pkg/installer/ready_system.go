package installer

import (
	"github.com/theprantadutta/kubernetes_install/pkg/logger"
	"github.com/theprantadutta/kubernetes_install/pkg/utils"
)

func ReadySystem() {
	log := logger.New()

	// Running System Update
	log.Info("Running System Upgrade...")
	utils.RunCommand("sudo apt-get update -y && sudo apt-get upgrade -y")
	log.Success("System Upgraded Successfully")

	// Disabling Swap
	log.Info("Disabling Swap...")
	utils.RunCommand("sudo swapoff -a")
	pattern := "/swap.img" // Pattern to search for in /etc/fstab
	if err := utils.CommentOutFstabLine(pattern); err != nil {
		log.Error("Error: %v\n", err)
	}
	log.Success("Successfully commented out the line in /etc/fstab")
	utils.RunCommand("sudo mount -a")
	utils.RunCommand("free -h")
	log.Success("Disabled Swap Successfully")
}
