package utils

import (
	"testing"
)

func TestNewNet(t *testing.T) {
	net := NewNet()
	id, err := net.NetworkId()
	if err != nil {
		t.Errorf("net networkId err:%v", err)
	}
	t.Logf("net network id:%v", id)
	isListening, err := net.Listening()
	if err != nil {
		t.Errorf("net listening err:%v", err)
	}
	t.Logf("net listening:%v", isListening)
	count, err := net.PeerCount()
	if err != nil {
		t.Errorf("net perrCount err:%v", err)
	}
	t.Logf("net peerCOunt:%v", count)
}
