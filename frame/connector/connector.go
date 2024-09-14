package connector

import (
	"common/logs"
	"fmt"
	"frame/game"
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
	return &Connector{
		handlers: make(net.LogicHandler),
	}
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
	err := c.remoteCli.Run()
	if err != nil {
		logs.Fatal("c.remoteCli.Run failed,err:%v", err.Error())
		return
	}
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
	logs.Info("run connector:%v", serverId)

	//游戏中的配置 读取 一般采用json的方式 需要读取json的配置文件
	c.wsManager.ServerId = serverId

	connectorConfig := game.Conf.GetConnector(serverId)
	if connectorConfig == nil {
		logs.Fatal("no connector config found")
		return
	}
	addr := fmt.Sprintf("%s:%d", connectorConfig.Host, connectorConfig.ClientPort)
	c.isRunning = true
	c.wsManager.Run(addr)

}

func (c *Connector) RegisterHandler(handlers net.LogicHandler) {
	c.handlers = handlers
}
