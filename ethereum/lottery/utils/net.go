package utils

import (
	"github.com/ethereum/go-ethereum/rpc"
	"log"
)

type Net struct {
	Client *rpc.Client
}

func NewNet() *Net {
	c, err := rpc.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("rpc dial err:%v", err)
	}
	return &Net{Client: c}
}

func (n *Net) NetworkId() (id string, err error) {
	err = n.Client.Call(&id, "net_version")
	if err != nil {
		return "", err
	}
	return
}

func (n *Net) Listening() (IsListening bool, err error) {
	err = n.Client.Call(&IsListening, "net_listening")
	if err != nil {
		return false, err
	}
	return
}

func (n *Net) PeerCount() (count string, err error) {
	err = n.Client.Call(&count, "net_peerCount")
	if err != nil {
		return "", err
	}
	return
}