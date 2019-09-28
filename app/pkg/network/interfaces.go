package network

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

var (
	netmaskPattern  = regexp.MustCompile("((\\d){1,3}\\.){3}(\\d){1,3}\\/(\\d){1,3}")
	IfaceNetmaskMap = make(map[string]string)
)

func Interfaces() error {

	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("Failed to get network Interfaces. Error: '%s'\n", err.Error())
	}

	for _, i := range ifaces {

		if !strings.Contains(i.Flags.String(), "up") {
			fmt.Printf("Interface '%s' is not UP\n", i.Name)
			continue
		}

		addrs, err := i.Addrs()
		if err != nil {
			fmt.Printf("Failed to fetch Addrs(). Error: '%s'\n", err.Error())
		}

		for _, a := range addrs {
			if netmaskPattern.MatchString(a.String()) {
				IfaceNetmaskMap[i.Name] = a.String()
			}
		}
	}

	for k, v := range IfaceNetmaskMap {
		fmt.Printf("Interface: '%s', Netmask: '%s'\n", k, v)
	}

	return nil
}
