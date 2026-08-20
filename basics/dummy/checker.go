package dummy

import (
	"fmt"
	"net"
	"time"
)

func CheckPort() {

	fmt.Println("just a call")

}

func CheckPort1(host string, port string) bool {
	CheckPort()
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
