package main

import (
	"basic-module/dummy"
	"fmt"
)

func main() {

	host := "google.com"
	port := "443"

	if dummy.CheckPort1(host, port) {
		fmt.Printf("%s:%s is OPEN\n", host, port)
	} else {
		fmt.Printf("%s:%s is CLOSED\n", host, port)
	}
}
