package netutil

import (
	"net"
	"testing"

	"lib/utils/internal"
)

func TestGetInternalIp(t *testing.T) {
	assert := internal.NewAssert(t, "TestGetInternalIp")

	internalIp := GetInternalIP()
	ip := net.ParseIP(internalIp)
	t.Logf("ip: %v", ip)
	assert.IsNotNil(ip)
}

func TestGetPublicIpInfo(t *testing.T) {
	assert := internal.NewAssert(t, "TestGetPublicIpInfo")

	publicIpInfo, err := GetPublicIPInfo()
	assert.IsNil(err)

	t.Logf("public ip info is: %+v \n", *publicIpInfo)
}

func TestIsPublicIP(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsPublicIP")

	ips := []net.IP{
		net.ParseIP("127.0.0.1"),
		net.ParseIP("192.168.0.1"),
		net.ParseIP("10.91.210.131"),
		net.ParseIP("172.20.16.1"),
		net.ParseIP("36.112.24.10"),
	}

	expected := []bool{false, false, false, false, true}

	for i := 0; i < len(ips); i++ {
		actual := IsPublicIP(ips[i])
		assert.Equal(expected[i], actual)
	}
}

func TestGetIps(t *testing.T) {
	ips := GetIPs()
	t.Log(ips)
}

func TestGetMacAddrs(t *testing.T) {
	macAddrs := GetMacAddrs()
	t.Log(macAddrs)
}
