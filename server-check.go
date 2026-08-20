package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkIP(ip string) string {

	parsedIP := net.ParseIP(ip)

	if parsedIP == nil {
		return "Invalid"
	}

	if parsedIP.To4() != nil {
		return "IPv4"
	}

	return "IPv6"
}

func checkPort(host string, port string) bool {

	address := net.JoinHostPort(host, port)

	conn, err := net.DialTimeout(
		"tcp",
		address,
		3*time.Second,
	)

	if err != nil {
		return false
	}

	defer conn.Close()

	return true
}

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run server-check.go <IP> <PORT>")
		return
	}

	host := os.Args[1]
	port := os.Args[2]

	ipType := checkIP(host)

	if ipType == "Invalid" {
		fmt.Println("Error: Invalid IP address")
		return
	}

	fmt.Println("================================")
	fmt.Println("       SERVER STATUS")
	fmt.Println("================================")
	fmt.Println()

	fmt.Println("Host   :", host)
	fmt.Println("Type   :", ipType)
	fmt.Println("Port   :", port)

	if checkPort(host, port) {
		fmt.Println("Status : UP")
	} else {
		fmt.Println("Status : DOWN")
	}

	fmt.Println()
	fmt.Println("================================")
}
