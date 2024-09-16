package installer

import (
	"fmt"
	"os"

	"github.com/theprantadutta/kubernetes_install/pkg/logger"
	"github.com/theprantadutta/kubernetes_install/pkg/utils"
)

func AddLoadBalancerIp() {
	log := logger.New()

	log.Info("Adding Load Balancer IP...")

	var loadBalancerIp string
	log.Info("Enter Load Balancer IP: ")
	fmt.Scan(&loadBalancerIp)
	if !utils.IsValidIP(loadBalancerIp) {
		log.Error("Failed to parse the IP")
		os.Exit(1)
	}
	log.Info("The provided IP is %s", loadBalancerIp)

	var loadBalancerAddr string
	log.Info("Enter Load Balancer Hostname: ")
	fmt.Scan(&loadBalancerAddr)
	log.Info("The provided Hostname is %s", loadBalancerAddr)

	if err := utils.AppendToHostsFile(loadBalancerIp, loadBalancerAddr); err != nil {
		log.Error("Error: %v\n", err)
	}
	log.Info("Successfully updated /etc/hosts")

	log.Success("Successfully Added Load Balaner IP")
}
