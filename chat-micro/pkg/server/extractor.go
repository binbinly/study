package server

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"chat-micro/pkg/registry"
	"chat-micro/pkg/util/addr"
)

func serviceDef(opts Options) *registry.Service {
	var port int
	var host string

	parts := strings.Split(opts.Address, ":")
	if len(parts) > 1 {
		host = strings.Join(parts[:len(parts)-1], ":")
		port, _ = strconv.Atoi(parts[len(parts)-1])
	} else {
		host = parts[0]
	}

	address, err := addr.Extract(host)
	if err != nil {
		address = host
	}

	node := &registry.Node{
		Id:       opts.Name + "-" + opts.Id,
		Address:  net.JoinHostPort(address, fmt.Sprint(port)),
		Metadata: opts.Metadata,
	}

	node.Metadata["registry"] = opts.Registry.String()

	return &registry.Service{
		Name:    opts.Name,
		Version: opts.Version,
		Nodes:   []*registry.Node{node},
	}
}
