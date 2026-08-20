package main

import (
	"fmt"
	"net"
)

func main() {
	ip := "172.16.20.204"

	parsedIP := net.ParseIP(ip)

	if parsedIP == nil {
		fmt.Println("Invalid IP address")
		return
	}

	if parsedIP.To4() != nil {
		fmt.Println("IPv4")
	} else {
		fmt.Println("IPv6")
	}
}
