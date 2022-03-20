package utils

import (
	"github.com/nacos-group/nacos-sdk-go/common/logger"
	"net"
)

var (
	localIP     = ""
	privateCIDR []*net.IPNet
)

// LocalIPv4 get local host IPv4 address
func LocalIPv4() string {
	if localIP != "" {
		return localIP
	}

	faces, err := getFaces()
	if err != nil {
		return ""
	}

	for _, address := range faces {
		ipNet, ok := address.(*net.IPNet)
		if !ok || ipNet.IP.To4() == nil || isFilteredIP(ipNet.IP) {
			continue
		}

		localIP = ipNet.IP.String()
		break
	}

	if localIP != "" {
		logger.Infof("Local IP:%s", localIP)
	}

	return localIP
}

func isFilteredIP(ip net.IP) bool {
	for _, privateIP := range privateCIDR {
		if privateIP.Contains(ip) {
			return true
		}
	}
	return false
}

// getFaces return addresses from interfaces that is up
func getFaces() ([]net.Addr, error) {
	var upAddrs []net.Addr

	interfaces, err := net.Interfaces()
	if err != nil {
		logger.Errorf("get Interfaces failed,err:%+v", err)
		return nil, err
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		if (iface.Flags & net.FlagLoopback) != 0 {
			continue
		}

		addresses, err := iface.Addrs()
		if err != nil {
			logger.Errorf("get InterfaceAddress failed,err:%+v", err)
			return nil, err
		}

		upAddrs = append(upAddrs, addresses...)
	}

	return upAddrs, nil
}
