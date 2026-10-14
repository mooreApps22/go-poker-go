package client

import (
	"fmt"
	"net"
)

func CreateDial(host string, port int) (net.Conn, error) {
	address := fmt.Sprintf("%s:%d", host, port)

	return net.Dial("tcp", address)
}
