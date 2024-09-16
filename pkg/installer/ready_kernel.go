package installer

import (
	"github.com/theprantadutta/kubernetes_install/pkg/logger"
	"github.com/theprantadutta/kubernetes_install/pkg/utils"
)

func ReadyKernel() {
	log := logger.New()

	log.Info("Making Kernel Ready, Please Wait...")

	// Load kernel modules
	utils.RunCommand(`sudo tee /etc/modules-load.d/containerd.conf <<EOF
overlay
br_netfilter
EOF`)

	utils.RunCommand("sudo modprobe overlay")
	utils.RunCommand("sudo modprobe br_netfilter")

	// Set kernel parameters for Kubernetes
	utils.RunCommand(`sudo tee /etc/sysctl.d/kubernetes.conf <<EOF
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
EOF`)

	// Reload sysctl
	utils.RunCommand("sudo sysctl --system")

	log.Success("Kernel is now ready to install containerd")
}
