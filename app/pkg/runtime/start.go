package runtime

import (
	"fmt"
	"log"
	"net"

	"github.com/ff/app/pkg/k8s"
	"github.com/ff/app/pkg/network"
)

// Start figures out the interface that needs to be monitored for network
// capture and initiates the network capture
func Start() {

	// Saves the count of IPs connected on an interface
	IfaceCount := make(map[string]int)

	// Populate the IfaceCount map
	for ip, labels := range k8s.IPPodMap {

		fmt.Printf("IPAddress: '%s', Labels: '%s'", ip, labels)

		for iface, cidr := range network.IfaceCIDRMap {

			// Parse CIDR
			_, ip4Net, err := net.ParseCIDR(cidr)
			if err != nil {
				log.Printf("Failed to parse CIDR '%s'. Error: '%s'", cidr, err.Error())
				continue
			}

			// Check if the "ip" address belongs to the current "cidr"
			if ip4Net.Contains(net.ParseIP(ip)) {
				fmt.Printf("ip '%s' belongs to network '%s' on interface '%s'", ip, cidr, iface)
				IfaceCount[iface]++
			}
		}
	}

	fmt.Printf("IfaceCount:\n'%+v'\n", IfaceCount)

	// Figure out the network interface to be used for network capture
	var iface string
	count := 0

	for k, v := range IfaceCount {
		if v > count {
			iface = k
			count = v
		}
	}

	log.Printf("Monitoring interface: '%s'", iface)

	// Initiate network capture
	network.Capture(iface)
}