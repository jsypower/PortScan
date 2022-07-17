package scanner

import (
	"net"
	"time"
)

func SynScan(protocol string, host string, port string) bool {
	address := host + ":" + port
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		return false
	}
	defer conn.Close()

	return true
}
