package global

import (
	"blog_server/config"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	CONFIG *config.Config
	DB     *gorm.DB
	LOG    *logrus.Logger
)
