package config

import "strconv"

type Mysql struct {
	Host         string `yaml: "host"`
	Port         int    `yaml: "port"`
	Db           string `yaml: "db"`
	User         string `yaml: "user"`
	Password     string `yaml: "password"`
	DBConfig     string `yaml:"dbconfig"`     //db高级配置
	MaxIdleConns int    `yaml:"maxidleconns"` //最大空闲连接数
	MaxOpenConns int    `yaml:"maxopenconns"` //最大连接容纳数
}

func (m Mysql) DSN() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.Db + "?" + m.DBConfig
}
