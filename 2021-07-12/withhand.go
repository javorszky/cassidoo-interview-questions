package july122021

import (
	"encoding/binary"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var reCIDR = regexp.MustCompile(`^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})/(\d{1,2})`)
var reIP = regexp.MustCompile(`^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})$`)

// cidr is a hand written struct that does not rely on the net package's cidr / ip parsing.
type cidr struct {
	notation    string
	baseIP      []byte
	networkMask byte
}

// newCIDR is a constructor returning a hand written cidr struct after some parsing.
func newCIDR(incoming string) (cidr, error) {
	matches := reCIDR.FindStringSubmatch(incoming)
	if len(matches) < 6 {
		return cidr{}, errors.New("incoming cidr string could not be parsed to cidr data")
	}

	ips, err := parseIPs(matches[1:5])
	if err != nil {
		return cidr{}, fmt.Errorf("failed to parse ip part of cidr: %w", err)
	}
	// parse cidr part
	mask, err := strconv.Atoi(matches[5])
	if err != nil {
		return cidr{}, fmt.Errorf("failed to parse cidr mask: %w", err)
	}
	if mask < 0 || mask > 32 {
		return cidr{}, fmt.Errorf("cidr mask is out of range. Should be 0-32, got %d", mask)
	}

	return cidr{
		notation:    incoming,
		baseIP:      ips,
		networkMask: byte(mask),
	}, nil
}

// InRange is a method on the hand written cidr struct to figure out whether a passed IP is in range. Does not use the
// built in net package to help.
func (c cidr) InRange(ip string) (bool, error) {
	ipParts := reIP.FindStringSubmatch(ip)
	if len(ipParts) < 5 {
		return false, fmt.Errorf("failed to parse ip string into parts: length should be 5, got %d", len(ipParts))
	}

	ips, err := parseIPs(ipParts[1:])
	if err != nil {
		return false, fmt.Errorf("failed to parse ip parts into ips: %w", err)
	}

	base := binary.BigEndian.Uint32(c.baseIP)
	inc := binary.BigEndian.Uint32(ips)

	xored := base ^ inc

	var computedMask uint32 = 2 << (31 - c.networkMask)

	return xored < computedMask, nil
}

// parseIPs is a utility function. Takes a slice of strings, turns them into a slice of bytes after checking some very
// basic rules. The first value needs to be larger than 0, otherwise all values should be 0,255 inclusive.
func parseIPs(parts []string) ([]byte, error) {
	ips := make([]byte, 0, 4)

	// parse first part of ip address
	ip1, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("failed to parse first ip bit: %w", err)
	}
	if ip1 < 1 || ip1 > 255 {
		return nil, fmt.Errorf("ip1 is out of range. Should be 1-255, got %d", ip1)
	}

	ips = append(ips, byte(ip1))

	for _, i := range parts[1:] {
		ip, err := strconv.Atoi(i)
		if err != nil {
			return nil, fmt.Errorf("faild to parse ip: %w", err)
		}
		if ip < 0 || ip > 255 {
			return nil, fmt.Errorf("ip is out of range. Should be 0-255, got %d", ip)
		}
		ips = append(ips, byte(ip))
	}
	return ips, nil
}

// inRange is the function to call with the input as specified by the newsletter.
func inRange(ip, cidr string) (bool, error) {
	c, err := newCIDR(cidr)
	if err != nil {
		return false, fmt.Errorf("got an error while parsing cidr: %s", err)
	}

	in, err := c.InRange(ip)
	if err != nil {
		return false, fmt.Errorf("got an error from inrange: %s", err)
	}
	return in, nil
}
