package installer

import (
	"github.com/theprantadutta/kubernetes_install/pkg/logger"
	"github.com/theprantadutta/kubernetes_install/pkg/utils"
)

// InstallContainerd installs and configures containerd runtime
func InstallContainerd() {
	log := logger.New()

	log.Info("Installing Containered in the system...")

	// Install dependencies and containerd
	utils.RunCommand(`sudo apt update`)
	utils.RunCommand(`sudo apt install -y curl gnupg2 software-properties-common apt-transport-https ca-certificates`)
	utils.RunCommand(`sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmour -o /etc/apt/trusted.gpg.d/docker.gpg`)
	utils.RunCommand(`sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"`)
	utils.RunCommand(`sudo apt update`)
	utils.RunCommand(`sudo apt install -y containerd.io`)

	// Configure containerd
	utils.RunCommand(`containerd config default | sudo tee /etc/containerd/config.toml >/dev/null 2>&1`)
	utils.RunCommand(`sudo sed -i 's/SystemdCgroup = false/SystemdCgroup = true/g' /etc/containerd/config.toml`)
	utils.RunCommand(`sudo systemctl restart containerd`)
	utils.RunCommand(`sudo systemctl enable containerd`)

	log.Success("Successfully installed containered in the system")
}
