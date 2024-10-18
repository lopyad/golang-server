package network

import (
	"github.com/gin-gonic/gin"
	"golang-server/service"
)

type Network struct {
	engine  *gin.Engine
	service *service.Service
}

func NewNetwork(service *service.Service) *Network {
	r := &Network{
		engine:  gin.New(),
		service: service,
	}
	newUserRouter(r, service.User)

	return r
}

func (n *Network) ServerStart(port string) error {
	port = ":" + port
	return n.engine.Run(port)
}
