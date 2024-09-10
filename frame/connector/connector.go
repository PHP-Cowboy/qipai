package connector

import (
	"frame/net"
	"frame/remote"
)

type Connector struct {
	isRunning bool
	wsManager *net.WsManager
	handlers  net.LogicHandler
	remoteCli remote.Client
}

func NewConnector() *Connector {
	return &Connector{}
}

func (c Connector) Run(serverId string) {
	if c.isRunning {
		return
	}

	//启动websocket和nats
	c.wsManager = net.NewWsManager()
	c.wsManager.ConnectorHandlers = c.handlers
	//启动nats nats server不会存储消息
	c.remoteCli = remote.NewNatsClient(serverId, c.wsManager.RemoteReadChan)
	c.remoteCli.Run()
	c.wsManager.RemoteCli = c.remoteCli
	c.Serve(serverId)
}

func (c *Connector) Close() {
	if !c.isRunning {
		return
	}
	c.wsManager.Close()
}

func (c *Connector) Serve(serverId string) {
	c.wsManager.ServerId = serverId

	c.wsManager.Run()
}

func (c *Connector) RegisterHandler(handlers net.LogicHandler) {
	c.handlers = handlers
}
