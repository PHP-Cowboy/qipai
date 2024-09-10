package router

import (
	"connector/handler"
	"core/repo"
	"frame/net"
)

func Register(r *repo.RepoManager) net.LogicHandler {
	handlers := make(net.LogicHandler)
	entryHandler := handler.NewEntryHandler(r)
	handlers["entryHandler.entry"] = entryHandler.Enter

	return handlers
}
