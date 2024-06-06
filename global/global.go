package global

import (
	"blog_server/config"

	"gorm.io/gorm"
)

var (
	CONFIG *config.Config
	DB     *gorm.DB
)
