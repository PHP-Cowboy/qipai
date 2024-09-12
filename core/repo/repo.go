package repo

import "common/database"

type RepoManager struct {
	Redis *database.RedisManager
	Db    *database.Db
}

func (m *RepoManager) Close() {

	if m.Redis != nil {
		m.Redis.Close()
	}
}

func NewRepoManager() *RepoManager {
	return &RepoManager{
		Redis: database.NewRedis(),
		Db:    database.NewDb(),
	}
}
