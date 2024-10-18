package cmd

import (
	"golang-server/config"
	"golang-server/network"
	"golang-server/repository"
	"golang-server/service"
)

type Cmd struct {
	config *config.Config

	network    *network.Network
	service    *service.Service
	repository *repository.Repository
}

func NewCmd() *Cmd {

	c := &Cmd{
		config: config.NewConfig(),
	}
	c.repository = repository.NewRepository()
	c.service = service.NewService(c.repository)
	c.network = network.NewNetwork(c.service)
	c.network.ServerStart(c.config.Server.Port)

	return c
}
