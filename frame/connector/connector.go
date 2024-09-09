package connector

import "frame/net"

type Connector struct {
	isRunning bool
	wsManager net.WsManager
}

func NewConnector() *Connector {
	return &Connector{}
}

func (c Connector) Run() {
	if c.isRunning {
		return
	}
}

func (c *Connector) Stop() {
	if !c.isRunning {
		return
	}
}

func (c *Connector) Serve() {
	c.wsManager.Run()
}
