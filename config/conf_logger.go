package config

type Logger struct {
	Level     string `yaml:"level"`     //日志等级
	Prefix    string `yaml:"prefix"`    //日志前缀
	Show_line bool   `yaml:"show_line"` //是否显示行号
}
