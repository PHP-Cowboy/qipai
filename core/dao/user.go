package dao

import "core/repo"

type UserDao struct {
	m *repo.RepoManager
}

func NewUserDao(m *repo.RepoManager) *UserDao {
	return &UserDao{
		m: m,
	}
}
