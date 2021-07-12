package july122021

import (
	"errors"
	"fmt"
	"net"
)

func inRangeWithNet(ipString, cidr string) (bool, error) {
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return false, fmt.Errorf("failed to parse cidr with net: %w", err)
	}

	incomingIP := net.ParseIP(ipString)
	if incomingIP == nil {
		return false, errors.New("failed to parse incoming ip address")
	}

	return ipNet.Contains(incomingIP), nil
}
