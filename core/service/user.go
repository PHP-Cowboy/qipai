package service

import (
	"core/dao"
	"core/repo"
)

type UserService struct {
	userDao *dao.UserDao
}

func NewUserService(r *repo.RepoManager) *UserService {
	return &UserService{
		userDao: dao.NewUserDao(r),
	}
}
