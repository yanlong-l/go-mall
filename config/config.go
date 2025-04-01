package config

import "time"

var (
	App      *appConfig
	Database *databaseConfig
)

type appConfig struct {
	Env  string `mapstructure:"env"`
	Name string `mapstructure:"name"`
	Log  struct {
		FilePath         string `mapstructure:"path"`
		FileMaxSize      int    `mapstructure:"max_size"`
		BackUpFileMaxAge int    `mapstructure:"max_age"`
	}
	Pagination struct {
		DefaultSize int `mapstructure:"default_size"`
		MaxSize     int `mapstructure:"max_size"`
	}
}

type databaseConfig struct {
	Type   string          `mapstructure:"type"`
	Master DbConnectOption `mapstructure:"master"`
	Slave  DbConnectOption `mapstructure:"slave"`
}

type DbConnectOption struct {
	DSN         string        `mapstructure:"dsn"`
	MaxOpenConn int           `mapstructure:"maxopen"`
	MaxIdleConn int           `mapstructure:"maxidle"`
	MaxLifeTime time.Duration `mapstructure:"maxlifetime"`
}
