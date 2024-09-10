package repo

import "common/database"

type RepoManager struct {
	Redis *database.RedisManager
}

func (m *RepoManager) Close() {

	if m.Redis != nil {
		m.Redis.Close()
	}
}

func NewRepoManager() *RepoManager {
	return &RepoManager{
		Redis: database.NewRedis(),
	}
}
