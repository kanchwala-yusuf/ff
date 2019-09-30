package main

import (
	"fmt"
	"net"

	"github.com/ff/app/pkg/k8s"
	"github.com/ff/app/pkg/network"
)

func main() {

	// Get namespaces
	namespaces, err := k8s.GetAllNamespaces()
	if err != nil {
		fmt.Printf("GetAllNamespaces() error: '%s'", err.Error())
	}

	for _, ns := range namespaces {

		/*
		if ns == "kube-system" {
			continue
		}
		*/

		err := k8s.GetPods(ns)
		if err != nil {
			fmt.Printf("GetPods() error: '%s'", err.Error())
		}
	}

	if err = network.Interfaces(); err != nil {
		fmt.Printf("Failed to fetch network interfaces. Error: '%s'", err.Error())
	}

	IfMap := make(map[string]int)

	for ip, labels := range k8s.IPCache {
		fmt.Printf("IPAddress: '%s', Labels: '%s'\n", ip, labels)

		for iface, netmask := range network.IfaceNetmaskMap {
			_, ip4Net, err := net.ParseCIDR(netmask)
			if err != nil {
				fmt.Printf("Failed to parse CIDR '%s'. Error: '%s'", netmask, err.Error())
			}

			if ip4Net.Contains(net.ParseIP(ip)) {
				fmt.Printf("ip '%s' belongs to network '%s' on interface '%s'\n", ip, netmask, iface)
				IfMap[iface] += 1
			}
		}
	}

	fmt.Printf("IfMap:\n'%+v'\n", IfMap)

	var iface string
	count := 0

	for k, v := range IfMap {
		if v > count {
			iface = k
			count = v
		}
	}

	fmt.Printf("Monitor interface: '%s'\n", iface)

	network.NetworkCapture(iface)
}
