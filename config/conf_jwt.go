package config

type JWT struct {
	DeadTime   int    `yaml:"deadtime"`   //多少秒过期
	SigningKey string `yaml:"signingkey"` //服务器加密密钥
}
