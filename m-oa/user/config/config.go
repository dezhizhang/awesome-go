package config

import (
	"common/logs"
	"github.com/spf13/viper"
	"log"
	"os"
)

var C = InitConfig()

type Config struct {
	viper *viper.Viper
	SC    *ServerConfig
}

type ServerConfig struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
}

func InitConfig() *Config {
	dir, _ := os.Getwd()
	conf := &Config{viper: viper.New()}

	conf.viper.SetConfigName("config")
	conf.viper.SetConfigType("yaml")
	conf.viper.AddConfigPath(dir + "/config")

	err := conf.viper.ReadInConfig()

	if err != nil {
		log.Fatalln(err)
	}
	conf.ReadServerConfig()
	conf.InitZapLog()
	return conf

}

func (c *Config) ReadServerConfig() {
	sc := &ServerConfig{}
	sc.Name = c.viper.GetString("server.name")
	sc.Addr = c.viper.GetString("server.addr")
	c.SC = sc
}

func (c *Config) InitZapLog() {
	lc := &logs.LogConfig{
		MaxAge:        c.viper.GetInt("maxAge"),
		MaxSize:       c.viper.GetInt("maxSize"),
		MaxBackups:    c.viper.GetInt("maxBackups"),
		DebugFileName: c.viper.GetString("zap.debugFileName"),
		InfoFileName:  c.viper.GetString("zap.infoFileName"),
		WarnFileName:  c.viper.GetString("zap.warnFileName"),
	}

	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
}
