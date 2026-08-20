package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run ip-cli.go <IP>")
		return
	}

	ip := os.Args[1]

	parsedIP := net.ParseIP(ip)

	if parsedIP == nil {
		fmt.Println("Invalid IP address")
		return
	}

	fmt.Println("IP:", ip)

	if parsedIP.To4() != nil {
		fmt.Println("Type: IPv4")
	} else {
		fmt.Println("Type: IPv6")
	}
}
