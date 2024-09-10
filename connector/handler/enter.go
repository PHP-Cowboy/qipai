package handler

import (
	"core/repo"
	"core/service"
	"frame/net"
)

type EnterHandler struct {
	userService *service.UserService
}

func (h *EnterHandler) Enter(session *net.Session, body []byte) (any, error) {
	return nil, nil
}

func NewEntryHandler(r *repo.RepoManager) *EnterHandler {
	return &EnterHandler{
		userService: service.NewUserService(r),
	}
}
