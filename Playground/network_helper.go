// networkhelper - Generates the range of IP Addresses, by taking NetworkIP and the Netmask.
// Also, gives you the Usable IP Addresses, First Usable, Last Usable, Gateway, Broadcast IP Addresses

package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
)

// parseArgs - To parse the command line arguments and verify them.
// Returns the IP Address and CIDR notation of Netmask.
func parseArgs() (string, string) {
	flagset := flag.NewFlagSet("flagset", flag.ExitOnError)
	networkip := flagset.String("ip", "", "A Valid Network IP")
	netmask := flagset.String("mask", "", "A Valid Network Mask")
	err := flagset.Parse(os.Args[1:])
	if err == nil {
		ipaddr := net.ParseIP(*networkip)
		nmask := net.ParseIP(*netmask)
		if ipaddr != nil && nmask != nil {
			//Converting nmask to CIDR format
			netmasksize, _ := net.IPMask(nmask.To4()).Size()
			netmaskcidr := strconv.Itoa(netmasksize)
			return ipaddr.String(), netmaskcidr
		}
		return "", ""
	}
	fmt.Println("Invalid Arguments! Exiting...")
	fmt.Println("Run `" + string(os.Args[0]) + " --help` for more details")
	return "", ""
}

//Hosts - To generate the IP ranges
func Hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}
	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); ipinc(ip) {
		ips = append(ips, ip.String())
	}
	return ips, nil
}

// ipinc - Increments the IP
func ipinc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// main - Main func
func main() {
	networkip, netmaskcidr := parseArgs()
	if networkip == "" || netmaskcidr == "" {
		fmt.Println("Missing Arguments! Exiting...")
		os.Exit(1)
	}
	hosts, _ := Hosts(networkip + "/" + netmaskcidr)
	fmt.Printf("\nUsable Host Addresses: %d \n", len(hosts)-3)
	fmt.Printf("Network Address:       %s \n", hosts[0])
	fmt.Printf("Gateway Address:       %s \n", hosts[1])
	fmt.Printf("First Host Address:    %s \n", hosts[2])
	fmt.Printf("Last Host Address:     %s \n", hosts[(len(hosts)-2)])
	fmt.Printf("Broadcast Address:     %s \n\n", hosts[(len(hosts)-1)])
}
