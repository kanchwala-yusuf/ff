package main

import (
	"fmt"
	"net"
    "log"
)

func main() {

    netMask := "10.8.5.1/24"
    // netMask := "10.8.4.1/24"

    ip := "10.8.5.6"

	ipv4Addr, ipv4Net, err := net.ParseCIDR(netMask)
	if err != nil {
		log.Fatal(err)
	}
    fmt.Printf("ip: '%s'\n", ip)
	fmt.Printf("addr: %+v\nnet: %+v\n", ipv4Addr, ipv4Net)

	fmt.Printf("Contains: '%t'\n", ipv4Net.Contains(net.ParseIP(ip)))
}
