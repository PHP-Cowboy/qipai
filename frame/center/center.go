package center

import (
	"common/global"
	"fmt"
	"frame/game"
	"frame/net"
	"frame/remote"
)

type Center struct {
	isRunning bool
	wsManager *net.WsManager
	handlers  net.LogicHandler
	remoteCli remote.Client
}

func NewCenter() *Center {
	return &Center{
		handlers: make(net.LogicHandler),
	}
}

func (c Center) Run(serverId string) {
	if c.isRunning {
		return
	}

	//启动websocket和nats
	c.wsManager = net.NewWsManager()
	c.wsManager.CenterHandlers = c.handlers
	//启动nats nats server不会存储消息
	c.remoteCli = remote.NewNatsClient(serverId, c.wsManager.RemoteReadChan)
	err := c.remoteCli.Run()
	if err != nil {
		global.Logger["err"].Fatalf("c.remoteCli.Run failed,err:%v", err.Error())
		return
	}
	c.wsManager.RemoteCli = c.remoteCli
	c.Serve(serverId)
}

func (c *Center) Close() {
	if !c.isRunning {
		return
	}
	c.wsManager.Close()
}

func (c *Center) Serve(serverId string) {
	global.Logger["err"].Infof("run center:%v", serverId)

	//游戏中的配置 读取 一般采用json的方式 需要读取json的配置文件
	c.wsManager.ServerId = serverId

	cfg := game.Conf.GetCenter(serverId)
	if cfg == nil {
		global.Logger["err"].Fatal("no center config found")
		return
	}
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.ClientPort)
	c.isRunning = true
	c.wsManager.Run(addr)

}

func (c *Center) RegisterHandler(handlers net.LogicHandler) {
	c.handlers = handlers
}
