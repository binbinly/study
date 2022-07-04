package netutil

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
)

// PublicIPInfo public ip info: country, region, isp, city, lat, lon, ip
type PublicIPInfo struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Ip          string  `json:"query"`
}

// GetInternalIp return internal ipv4
func GetInternalIP() string {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		panic(err.Error())
	}
	for _, a := range addr {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}

	return ""
}

// GetPublicIPInfo return public ip information
// return the GetPublicIPInfo struct
func GetPublicIPInfo() (*PublicIPInfo, error) {
	resp, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ip PublicIPInfo
	err = json.Unmarshal(body, &ip)
	if err != nil {
		return nil, err
	}

	return &ip, nil
}

// GetIPs return all ipv4 of system
func GetIPs() []string {
	var ips []string

	adds, err := net.InterfaceAddrs()
	if err != nil {
		return ips
	}

	for _, addr := range adds {
		ipNet, isValid := addr.(*net.IPNet)
		if isValid && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}

	return ips
}

// GetMacAddrs get mac address
func GetMacAddrs() []string {
	var adds []string

	nets, err := net.Interfaces()
	if err != nil {
		return adds
	}

	for _, n := range nets {
		addr := n.HardwareAddr.String()
		if len(addr) == 0 {
			continue
		}
		adds = append(adds, addr)
	}

	return adds
}

// IsPublicIP verify a ip is public or not
func IsPublicIP(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsLinkLocalMulticast() || ip.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := ip.To4(); ip4 != nil {
		switch {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}