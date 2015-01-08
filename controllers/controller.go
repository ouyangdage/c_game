package controllers

import (
	"github.com/fhbzyc/c_game/network"
	"github.com/fhbzyc/c_game/protocol"
)

type Controller struct {
	Connect *network.Connect
	Request *protocol.Request
}
