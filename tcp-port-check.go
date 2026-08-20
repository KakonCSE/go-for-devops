package goprojects
package main

import (
	"fmt"
	"net"
	"time"
)

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

	conn.Close()

	return true
}

func main() {

	host := "google.com"
	port := "443"

	if checkPort(host, port) {
		fmt.Printf("%s:%s is OPEN\n", host, port)
	} else {
		fmt.Printf("%s:%s is CLOSED\n", host, port)
	}
}