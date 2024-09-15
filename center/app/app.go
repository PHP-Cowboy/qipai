package app

import (
	"center/router"
	"context"
	"core/repo"
	"frame/center"
	"log"
	"os"
	"os/signal"
)

func Run(ctx context.Context, serverId string) error {

	c := center.NewCenter()

	repoManager := repo.NewRepoManager()

	c.RegisterHandler(router.Register(repoManager))

	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)

	go func() {
		c.Run(serverId)
	}()

	<-quit
	log.Println("server is shutting down...")

	c.Close()

	close(done)
	log.Println("server stopped")

	return nil
}
