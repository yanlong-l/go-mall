package dao

import (
	"github.com/yanlong-l/go-mall/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB(option config.DbConnectOption) *gorm.DB {
	db, err := gorm.Open(mysql.Open(option.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(option.MaxIdleConn)
	sqlDB.SetMaxIdleConns(option.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(option.MaxLifeTime)
	if err = sqlDB.Ping(); err != nil {
		panic(err)
	}
	return db
}
