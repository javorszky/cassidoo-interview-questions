package july122021

import (
	"fmt"
	"log"
)

// Tasks is the entry point for today's code challenge.
func Tasks() {
	ipCIDRPairs := [][]string{
		{"192.168.4.27", "192.168.0.0/16"},
		{"192.168.4.27", "192.168.1.0/24"},
	}

	for _, pairs := range ipCIDRPairs {
		in, err := inRange(pairs[0], pairs[1])
		if err != nil {
			log.Fatalf("error from withhand inrange for [%s]/[%s]: %s", pairs[0], pairs[1], err)
		}
		fmt.Printf("withhand: ip [%s] in range of cidr [%s]: %t\n", pairs[0], pairs[1], in)

		inNet, err := inRangeWithNet(pairs[0], pairs[1])
		if err != nil {
			log.Fatalf("error from withnet inrange for [%s]/[%s]: %s", pairs[0], pairs[1], err)
		}
		fmt.Printf("withnet:  ip [%s] in range of cidr [%s]: %t\n\n", pairs[0], pairs[1], inNet)
	}
}
