package installer

import (
	"github.com/theprantadutta/kubernetes_install/pkg/logger"
	"github.com/theprantadutta/kubernetes_install/pkg/utils"
)

// InstallKubernetes installs Kubernetes components
func InstallKubernetes() {
	log := logger.New()

	log.Info("Installing Kubernetes with kubelet, kubeadm and kubectl in the system...")

	// Update the apt package index and install necessary packages
	utils.RunCommand(`sudo apt-get update`)
	utils.RunCommand(`sudo apt-get install -y apt-transport-https ca-certificates curl gnupg`)

	// Download the public signing key for the Kubernetes package repositories
	utils.RunCommand(`sudo mkdir -p -m 755 /etc/apt/keyrings`)
	utils.RunCommand(`curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.31/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg`)
	utils.RunCommand(`sudo chmod 644 /etc/apt/keyrings/kubernetes-apt-keyring.gpg`)

	// Add the Kubernetes apt repository
	utils.RunCommand(`echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.31/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list`)
	utils.RunCommand(`sudo chmod 644 /etc/apt/sources.list.d/kubernetes.list`)

	// Update apt package index and install Kubernetes components
	utils.RunCommand(`sudo apt-get update`)
	utils.RunCommand(`sudo apt-get install -y kubelet kubeadm kubectl`)
	utils.RunCommand(`sudo apt-mark hold kubelet kubeadm kubectl`)

	log.Success("Successfully installed Kubernetes with kubelet, kubeadm and kubectl in the system")
}
