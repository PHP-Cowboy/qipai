package game

import (
	"common/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
)

var Conf *Config

const (
	gameConfig = "gameConfig.json"
	servers    = "servers.json"
)

type Config struct {
	GameConfig  map[string]GameConfigValue `json:"gameConfig"`
	ServersConf ServersConf                `json:"serversConf"`
}
type ServersConf struct {
	Nats       NatsConfig       `json:"nats" `
	Center     []*CenterConfig  `json:"center" `
	Servers    []*ServersConfig `json:"servers" `
	TypeServer map[string][]*ServersConfig
}

type ServersConfig struct {
	ID               string `json:"id" `
	ServerType       string `json:"serverType" `
	HandleTimeOut    int    `json:"handleTimeOut" `
	RPCTimeOut       int    `json:"rpcTimeOut" `
	MaxRunRoutineNum int    `json:"maxRunRoutineNum" `
}

type CenterConfig struct {
	ID         string `json:"id" `
	Host       string `json:"host" `
	ClientPort int    `json:"clientPort" `
	Frontend   bool   `json:"frontend" `
	ServerType string `json:"serverType" `
}
type NatsConfig struct {
	Url string `json:"url"`
}

type GameConfigValue map[string]any

func InitConfig(configDir string) {
	Conf = new(Config)
	dir, err := os.ReadDir(configDir)
	if err != nil {
		global.Logger["err"].Fatalf("read config dir err:%v", err)
		return
	}
	for _, v := range dir {
		configFile := path.Join(configDir, v.Name())
		if v.Name() == gameConfig {
			readGameConfig(configFile)
		}
		if v.Name() == servers {
			readServersConfig(configFile)
		}
	}
}

func readServersConfig(configFile string) {
	var serversConfig ServersConf
	v := viper.New()
	v.SetConfigFile(configFile)
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		log.Println("serversConfig 配置文件被修改了")
		err := v.Unmarshal(&serversConfig)
		if err != nil {
			panic(fmt.Errorf("serversConfig Unmarshal change config data,err:%v \n", err))
		}
		Conf.ServersConf = serversConfig
	})
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("serversConfig 读取配置文件出错,err:%v \n", err))
	}
	//解析
	err = v.Unmarshal(&serversConfig)
	if err != nil {
		panic(fmt.Errorf("serversConfig Unmarshal config data,err:%v \n", err))
	}
	Conf.ServersConf = serversConfig
	typeServersConfig()
}

func typeServersConfig() {
	if len(Conf.ServersConf.Servers) > 0 {
		if Conf.ServersConf.TypeServer == nil {
			Conf.ServersConf.TypeServer = make(map[string][]*ServersConfig)
		}
		for _, v := range Conf.ServersConf.Servers {
			if Conf.ServersConf.TypeServer[v.ServerType] == nil {
				Conf.ServersConf.TypeServer[v.ServerType] = make([]*ServersConfig, 0)
			}
			Conf.ServersConf.TypeServer[v.ServerType] = append(Conf.ServersConf.TypeServer[v.ServerType], v)
		}
	}
}

func readGameConfig(configFile string) {
	gc := make(map[string]GameConfigValue)
	v := viper.New()
	v.SetConfigFile(configFile)
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		log.Println("gameConfig配置文件被修改了")
		err := v.Unmarshal(&gc)
		if err != nil {
			panic(fmt.Errorf("gameConfig Unmarshal change config data,err:%v \n", err))
		}
		Conf.GameConfig = gc
	})
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("gameConfig 读取配置文件出错,err:%v \n", err))
	}
	log.Println("%v", v.AllKeys())
	err = v.Unmarshal(&gc)
	if err != nil {
		panic(fmt.Errorf("gameConfig Unmarshal config data,err:%v \n", err))
	}
	Conf.GameConfig = gc

}

func (c *Config) GetCenter(serverId string) *CenterConfig {

	//b, _ := json.Marshal(c.ServersConf)
	//
	//fmt.Println(string(b))

	for _, v := range c.ServersConf.Center {
		if v.ID == serverId {
			return v
		}
	}
	return nil
}

func (c *Config) GetCenterByServerType(serverType string) *CenterConfig {
	for _, v := range c.ServersConf.Center {
		if v.ServerType == serverType {
			return v
		}
	}
	return nil
}

func (c *Config) GetFrontGameConfig() map[string]any {
	result := make(map[string]any)
	for k, v := range c.GameConfig {
		value, ok := v["value"]
		backend := false
		_, exist := v["backend"]
		if exist {
			backend = v["backend"].(bool)
		}
		if ok && !backend {
			result[k] = value
		}
	}
	return result
}
