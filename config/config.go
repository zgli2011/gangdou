package config

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// AppConfig : 配置文件映射
type AppConfig struct {
	Database struct {
		Adapter  string `yaml:"adapter"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DBName   string `yaml:"name"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

// InitConf : 初始化配置文件
func InitConf() *AppConfig {
	conf := new(AppConfig)
	config, err := ioutil.ReadFile("config/conf.yaml")
	if err != nil {
		fmt.Println("config file load failed: ", err)
	}
	err = yaml.Unmarshal(config, conf)
	if err != nil {
		fmt.Println("config file map failed: ", err)
	}
	return conf
}
