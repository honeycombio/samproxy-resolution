package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	var address string

	if len(os.Args) > 1 {
		address = os.Args[1]
	} else {
		// the default from the sample toml file
		address = "0.0.0.0:8080"
	}

	// this is the code that determines the value we store in redis for this shard
	port := strings.Split(address, ":")[1]
	hostname, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	public := fmt.Sprintf("http://%s:%s", hostname, port)

	fmt.Println("address added to redis:", public)

	// this is the code that is used to find "self" in the list of peers
	peerIPList, err := net.LookupHost(public)

	if err != nil {
		panic(err)
	}

	localAddrs, err := net.InterfaceAddrs()

	if err != nil {
		panic(err)
	}

	for _, peer := range peerIPList {
		fmt.Println("value from LookupHost:", peer)

		for _, local := range localAddrs {
			fmt.Println("value from InterfaceAddrs:", local)

			ip, _, err := net.ParseCIDR(local.String())

			if err != nil {
				panic(err)
			}

			fmt.Println("value from ParseCIDR:", ip)

			if peer == ip.String() {
				fmt.Println("Found match:", peer, ip.String())
			} else {
				fmt.Println("Failed to match:", peer, ip.String())
			}
		}
	}
}
