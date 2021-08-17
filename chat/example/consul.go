package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"sync"
)

var consulClient *api.Client
var wg sync.WaitGroup

func main() {
	config := api.DefaultConfig()
	config.Address = "192.168.8.76:8500"
	consulClient, _ = api.NewClient(config)

	wg.Add(1)
	go watcher()
	select {}
}

func watcher() {
	defer wg.Done()

	var lastIndex uint64
	for {
		services, m, err := consulClient.Health().Service("chat-logic", "", true, &api.QueryOptions{WaitIndex: lastIndex})
		//service, _, err := consul.Agent().Service("conn-1", nil)
		if err != nil {
			panic(err)
		}
		lastIndex = m.LastIndex
		fmt.Printf("service:%v\n", services)
	}
}
