package config

type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
}
type Mysql struct {
	Host     string `yaml: "host"`
	Port     int    `yaml: "port"`
	Db       string `yaml: "db"`
	User     string `yaml: "user"`
	Password string `yaml: "password"`
}
type Logger struct {
	Level     string `yaml:"level"`     //日志等级
	Prefix    string `yaml:"prefix"`    //日志前缀
	Show_line bool   `yaml:"show_line"` //是否显示行号
}
type System struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}
