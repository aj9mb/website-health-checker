package main

import (
	"fmt"
	"net"
	"time"
)

func Check(domain, port string, timeoutSec int) string {
	address := domain + ":" + port
	timeout := time.Duration(time.Duration(timeoutSec) * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)
	var status string
	if err != nil {
		status = fmt.Sprintf("[Down] %v is unreachable,\nError: %v", domain, err)
	} else {
		status = fmt.Sprintf("[Up] %v is reachable,\nFrom: %v\nTo: %v", domain, conn.LocalAddr(), conn.RemoteAddr())
		conn.Close()
	}
	return status
}
