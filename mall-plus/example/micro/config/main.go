package main

import (
	"fmt"
	"github.com/asim/go-micro/plugins/config/encoder/yaml/v4"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source/file"
)

type Host struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

var host Host

func main() {
	enc := yaml.NewEncoder()

	c, _ := config.NewConfig(
		config.WithReader(
			json.NewReader(
				reader.WithEncoder(enc))))

	if err := c.Load(file.NewSource(
		file.WithPath("./config.yaml"))); err != nil {
		fmt.Println(err)
		return
	}
	go watch(c)

	fmt.Println("data", c.Map())

	if err := c.Get("hosts", "database").Scan(&host); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(host.Address, host.Port)
	select {}
}

func watch(c config.Config) {
	for {
		w, err := c.Watch("hosts", "database")
		if err != nil {
			fmt.Println("err", err)
			return
		}
		v, err := w.Next()
		if err != nil {
			fmt.Printf("next err: %v\n", err)
			return
		}
		if err = v.Scan(&host); err != nil {
			fmt.Printf("scan err: %v\n", err)
			return
		}
		fmt.Println(host.Address, host.Port)
	}
}
