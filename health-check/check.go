package main

import (
	"fmt"
	"net"
	"time"
)

func Check(destination string, port string) string {
	address := destination + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)
	var status string
	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable, \n Error: %v", destination, err) // Printf prints formatted output to the console, while Sprintf returns the formatted output as a string.
	} else {
		status = fmt.Sprintf("[UP] %v is reachable, \n From: %v\n To: %v", destination, conn.LocalAddr(), conn.RemoteAddr())
		conn.Close() // Close the connection after the check
	}
	return status
}
