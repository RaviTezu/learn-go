package main

import (
	"fmt"
	"net"
	"time"
)

// HOST - You can define a FQDN or an IP here.
// PORT - Takes a port number as a string.
const (
	HOST = "127.0.0.1"
	PORT = "22"
)

// IsvalidIP - Checks whether a passed IP is valid or not.
func IsvalidIP() bool {
	ip := net.ParseIP(HOST)
	return ip != nil
}

// IsvalidFQDN - Checks for FQDN to IP resolution
func IsvalidFQDN() bool {
	ip, _ := net.LookupIP(HOST)
	return ip != nil
}

//Isalive - Takes in a host and port as strings and returns true if it is able to connect to that port
func Isalive() bool {
	conn, err := net.DialTimeout("tcp", HOST+":"+PORT, 3*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func main() {
	if IsvalidIP() || IsvalidFQDN() {
		fmt.Println("1")
	}
}
