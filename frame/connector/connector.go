package connector

import "frame/net"

type Connector struct {
	isRunning bool
	wsManager *net.WsManager
	handlers  net.LogicHandler
	clients   map[string]Connection
}

func NewConnector() *Connector {
	return &Connector{}
}

func (c Connector) Run(serverId string) {
	if c.isRunning {
		return
	}
}

func (c *Connector) Close() {
	if !c.isRunning {
		return
	}
	c.wsManager.Close()
}

func (c *Connector) Serve() {
	c.wsManager.Run()
}

func (c *Connector) RegisterHandler(handlers net.LogicHandler) {
	c.handlers = handlers
}
