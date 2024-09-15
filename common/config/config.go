package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf *Config

type Config struct {
	Database Database `json:"database"`
	Jwt      JwtConf  `json:"jwt"`
	Etcd     EtcdConf `json:"etcd"`
}

type JwtConf struct {
	Secret string `json:"secret"`
	Exp    int64  `json:"exp"`
}

// 数据库配置
type Database struct {
	Mysql MysqlConfig `json:"mysql"`
	Redis RedisConf   `json:"redis"`
}

type MysqlConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

type RedisConf struct {
	Addr            string   `json:"addr"`
	ClusterAddrList []string `json:"clusterAddrList"`
	Password        string   `json:"password"`
	PoolSize        int      `json:"poolSize"`
	MinIdleConnNum  int      `json:"minIdleConnNum"`
	Host            string   `json:"host"`
	Port            int      `json:"port"`
}
type EtcdConf struct {
	AddrList    []string       `json:"addrList"`
	RWTimeout   int            `json:"rwTimeout"`
	DialTimeout int            `json:"dialTimeout"`
	Register    RegisterServer `json:"register"`
}
type RegisterServer struct {
	Addr    string `json:"addr"`
	Name    string `json:"name"`
	Version string `json:"version"`
	Weight  int    `json:"weight"`
	Ttl     int64  `json:"ttl"`
}

// 加载配置
func InitConfig(confFile string) {
	Conf = new(Config)
	v := viper.New()
	v.SetConfigFile(confFile)
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		err := v.Unmarshal(&Conf)
		if err != nil {
			panic(fmt.Errorf("Unmarshal change config data,err:%v \n", err))
		}
	})
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件出错,err:%v \n", err))
	}
	//解析
	err = v.Unmarshal(&Conf)
	if err != nil {
		panic(fmt.Errorf("Unmarshal config data,err:%v \n", err))
	}
}
