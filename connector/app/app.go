package app

import (
	"common/config"
	"common/logs"
	"context"
	"frame/connector"
)

func Run(ctx context.Context) error {
	//1.做一个日志库 info error fatal debug
	logs.InitLog(config.Conf.AppName)

	go func() {
		c := connector.NewConnector()

		c.Run()
	}()

	return nil
}
