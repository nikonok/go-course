package main

import (
	"fmt" // output into terminal
	"log" // logging pkg
	"net" // main pkg for networking
)

func main() {
	// find ips of the domain
	ips, err := net.LookupIP("google.com")
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not get IPs: %v\n", err))
	}
	for _, ip := range ips {
		fmt.Printf("google.com. IN A %s\n", ip.String())
	}
}
