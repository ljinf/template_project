package config

import "time"

// 项目通过这里的变量读取应用配置中的对应项
var (
	App      *appConfig
	Database *databaseConfig
	Redis    *redisConfig
)

type appConfig struct {
	Name string `mapstructure:"name"`
	Env  string `mapstructure:"env"`
	Log  struct {
		LogLevel         string `mapstructure:"log_level"`
		FilePath         string `mapstructure:"path"`
		FileMaxSize      int    `mapstructure:"max_size"`
		BackUpFileMaxAge int    `mapstructure:"max_age"`
	}
	Pagination struct {
		DefaultSize int `mapstructure:"default_size"`
		MaxSize     int `mapstructure:"max_size"`
	} `mapstructure:"pagination"`
}

type databaseConfig struct {
	Type   string          `mapstructure:"type"`
	Master DbConnectOption `mapstructure:"master"`
	Slave  DbConnectOption `mapstructure:"slave"`
}

type DbConnectOption struct {
	DSN         string        `mapstructure:"dsn"`
	MaxOpenConn int           `mapstructure:"max_open"`
	MaxIdleConn int           `mapstructure:"max_idle"`
	MaxLifeTime time.Duration `mapstructure:"max_life_time"`
}

// Redis 配置
type redisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"pool_size"`
	DB       int    `mapstructure:"db"`
}
