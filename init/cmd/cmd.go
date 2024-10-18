package cmd

import (
	"golang-server/config"
	"golang-server/network"
)

type Cmd struct {
	config  *config.Config
	network *network.Network
}

func NewCmd() *Cmd {

	c := &Cmd{
		config:  config.NewConfig(),
		network: network.NewNetwork(),
	}
	
	c.network.ServerStart(c.config.Server.Port)
	return c
}
