package router

import (
	"center/handler"
	"core/repo"
	"frame/net"
)

func Register(r *repo.RepoManager) net.LogicHandler {
	handlers := make(net.LogicHandler)
	enterHandler := handler.NewEnterHandler(r)
	handlers["enterHandler.entry"] = enterHandler.Enter

	return handlers
}
