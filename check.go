package main

import (
	"net"
	"time"
)

func CheckHealth(domain string, port string) string {
	address := domain + ":" + port
	timeout := time.Duration(5 * time.Second)

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		println("DOWN from " + conn.LocalAddr().String() + " to " + conn.RemoteAddr().String())
		return "DOWN"
	}
	defer conn.Close()
	println("UP from " + conn.LocalAddr().String() + " to " + conn.RemoteAddr().String())
	return "UP"
}
