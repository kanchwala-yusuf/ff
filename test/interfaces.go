package network

import (
	"fmt"
	"net"
)

func Interfaces() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("Failed to get network interfaces. Error: '%s'\n", err.Error())
	}
	fmt.Printf("Interfaces:\n%+v\n", ifaces)

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Printf("Failed to fetch Addrs(). Error: '%s'\n", err.Error())
		}

		for _, a := range addrs {
			fmt.Printf("Name: '%s', AddrNetwork: '%s', AddrIP: '%s'\n", i.Name, a.Network(), a.String())
		}
	}
}
