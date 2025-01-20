package network

import (
	"fmt"
	"net"
)

func CalculateStartAndEndIpOfCIDRBlock(CIDR string) (string, string, error) {
	_, net, err := net.ParseCIDR(CIDR)
	if err != nil {
		fmt.Printf("Error parsing CIDR block: %s", CIDR)
		return "", "", err
	}

	startIP := net.IP.String()
	endIP := getEndIP(net)

	return startIP, endIP, nil
}

func getEndIP(cidrNet *net.IPNet) string {
	endIP := make(net.IP, len(cidrNet.IP))
	copy(endIP, cidrNet.IP)
	for i := 0; i < len(endIP); i++ {
		endIP[i] |= ^(cidrNet.Mask[i])
	}
	return endIP.String()
}
