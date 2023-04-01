package config

import (
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
	return conf

}

func (c *Config) ReadServerConfig() {
	sc := &ServerConfig{}
	sc.Name = c.viper.GetString("server.name")
	sc.Addr = c.viper.GetString("server.addr")
	c.SC = sc
}
