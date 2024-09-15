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
	//1.校验权限，获取用户信息，游戏房间信息
	//2.加入游戏，
	return "bbb", nil
}

func NewEnterHandler(r *repo.RepoManager) *EnterHandler {
	return &EnterHandler{
		userService: service.NewUserService(r),
	}
}
