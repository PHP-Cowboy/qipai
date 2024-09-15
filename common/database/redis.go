package database

import (
	"common/config"
	"common/global"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisManager struct {
	Cli        *redis.Client        //单机
	ClusterCli *redis.ClusterClient //集群
}

func NewRedis() *RedisManager {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var clusterCli *redis.ClusterClient
	var cli *redis.Client
	addrs := config.Conf.Database.Redis.ClusterAddrList
	if len(addrs) == 0 {
		//非集群 单节点
		cli = redis.NewClient(&redis.Options{
			Addr:         config.Conf.Database.Redis.Addr,
			PoolSize:     config.Conf.Database.Redis.PoolSize,
			MinIdleConns: config.Conf.Database.Redis.MinIdleConnNum,
			Password:     config.Conf.Database.Redis.Password,
		})
	} else {
		clusterCli = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        config.Conf.Database.Redis.ClusterAddrList,
			PoolSize:     config.Conf.Database.Redis.PoolSize,
			MinIdleConns: config.Conf.Database.Redis.MinIdleConnNum,
			Password:     config.Conf.Database.Redis.Password,
		})
	}
	if clusterCli != nil {
		if err := clusterCli.Ping(ctx).Err(); err != nil {
			global.Logger["err"].Fatalf("redis cluster connect err:%v", err)
			return nil
		}
	}
	if cli != nil {
		if err := cli.Ping(ctx).Err(); err != nil {
			global.Logger["err"].Fatalf("redis connect err:%v", err)
			return nil
		}
	}
	return &RedisManager{
		Cli:        cli,
		ClusterCli: clusterCli,
	}
}

func (r *RedisManager) Close() {
	if r.ClusterCli != nil {
		if err := r.ClusterCli.Close(); err != nil {
			global.Logger["err"].Errorf("redis cluster close err:%v", err)
		}
	}
	if r.Cli != nil {
		if err := r.Cli.Close(); err != nil {
			global.Logger["err"].Errorf("redis close err:%v", err)
		}
	}
}

func (r *RedisManager) Set(ctx context.Context, key, value string, expire time.Duration) error {
	if r.ClusterCli != nil {
		return r.ClusterCli.Set(ctx, key, value, expire).Err()
	}
	if r.Cli != nil {
		return r.Cli.Set(ctx, key, value, expire).Err()
	}
	return nil
}
