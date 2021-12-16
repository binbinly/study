package ip

import (
	"net"
	"strings"
	"sync"

	tnet "github.com/toolkits/net"
)

var (
	once     sync.Once
	clientIP = "127.0.0.1"
)

// GetLocalIP 获取本地内网IP
func GetLocalIP() string {
	once.Do(func() {
		ips, _ := tnet.IntranetIP()
		if len(ips) > 0 {
			clientIP = ips[0]
		} else {
			clientIP = "127.0.0.1"
		}
	})
	return clientIP
}

// GetInternalIP get internal ip.
func GetInternalIP() string {
	inters, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, inter := range inters {
		if inter.Flags&net.FlagUp != 0 && !strings.HasPrefix(inter.Name, "lo") {
			addressAll, err := inter.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addressAll {
				if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
					if ipNet.IP.To4() != nil {
						return ipNet.IP.String()
					}
				}
			}
		}
	}
	return clientIP
}
