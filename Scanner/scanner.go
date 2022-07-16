package scanner

import (
	"net"
	"strconv"
	"time"
)

func SynScan(protocol string, host string, port int) bool {
	address := host + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		return false
	}
	defer conn.Close()

	return true
}
