package server

import (
	"fmt"
	"net"
)

func CreateListener(port int) (net.Listener, error) {
	portString := fmt.Sprintf(":%d", port)

	return net.Listen("tcp", portString)
}
