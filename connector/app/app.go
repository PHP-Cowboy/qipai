package app

import (
	"common/config"
	"common/logs"
	"connector/router"
	"context"
	"core/repo"
	"frame/connector"
	"log"
	"os"
	"os/signal"
	"time"
)

func Run(ctx context.Context, serverId string) error {
	//1.做一个日志库 info error fatal debug
	logs.InitLog(config.Conf.AppName)

	c := connector.NewConnector()

	repoManager := repo.NewRepoManager()

	c.RegisterHandler(router.Register(repoManager))

	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)

	go func() {
		c.Run(serverId)
	}()

	<-quit
	log.Println("Websocket server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := c.Close(ctx); err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v", err)
	}
	close(done)
	log.Println("Websocket server stopped")

	return nil
}
