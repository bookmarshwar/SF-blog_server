package core

import (
	"blog_server/config"
	"blog_server/global"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// 配置读取
func InitConf() {
	const ConfigFile = "setting.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("配置文件错误: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("配置文件初始化:%s", err)
	}
	log.Println("配置文件初始化成功")
	//fmt.Println(c)
	global.CONFIG = c
}
