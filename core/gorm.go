package core

import (
	"blog_server/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {
	if global.CONFIG.Mysql.Host == "" {
		//LOG.没配置mysql,请检查配置文件是否正确
		global.LOG.Warnln("没配置mysql,请检查配置文件是否正确")
		return nil
	}
	dsn := global.CONFIG.Mysql.DSN()

	var mysqlLogger logger.Interface
	if global.CONFIG.System.Env == "dev" {
		//log.
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		//log.
		mysqlLogger = logger.Default.LogMode(logger.Error)

	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		//log.mysql连接失败
		global.LOG.Fatalf("[%s] mysql连接失败", dsn)
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(global.CONFIG.Mysql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(global.CONFIG.Mysql.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour * 4) //连接最大生命周期
	return db

}
