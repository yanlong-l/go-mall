package dao

import (
	"time"

	"github.com/yanlong-l/go-mall/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _DBMaster *gorm.DB
var _DBSlave *gorm.DB

func init() {
	_DBMaster = initDB(config.Database.Master)
	_DBSlave = initDB(config.Database.Slave)
}

func initDB(option config.DbConnectOption) *gorm.DB {
	db, err := gorm.Open(mysql.Open(
		option.DSN,
	), &gorm.Config{
		Logger: NewGormLogger(),
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(option.MaxOpenConn)
	sqlDB.SetMaxIdleConns(option.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(option.MaxLifeTime * time.Second)
	if err = sqlDB.Ping(); err != nil {
		panic(err)
	}
	return db
}

func DBMaster() *gorm.DB {
	return _DBMaster
}
func DBSlave() *gorm.DB {
	return _DBSlave
}
