package config

import (
	"github.com/BurntSushi/toml"
	"github.com/kataras/golog"
	"os"
	"reflect"
	"sync"
)

type EnvConfig struct {
	AppName  string      `toml:"app_name"`
	LogLevel string      `toml:"log_level"`
	Mysql    MysqlConifg `toml:"mysql"`
}

type MysqlConifg struct {
	DbHost           string `toml:"db_host"`
	DbUser           string `toml:"db_user"`
	DbDatabase       string `toml:"db_database"`
	DbPassword       string `toml:"db_password"`
	DbPort           int    `toml:"db_port"`
	DbCharset        string `toml:"db_charset"`
	ShowSql          bool   `toml:"show_sql"`
	PoolMaxConns     int    `toml:"pool_max_conns"`
	PoolMaxIdleConns int    `toml:"pool_max_idle_conns"`
}

var (
	EnvSettings EnvConfig
	lock        sync.Mutex
)

func (e EnvConfig) isEmpty() bool {
	return reflect.DeepEqual(e, EnvConfig{})
}

func GetEnv() *EnvConfig {
	if !EnvSettings.isEmpty() {
		return &EnvSettings
	}
	var envPath = "./.env.toml"

	_, statErr := os.Stat(envPath)
	if statErr != nil {
		golog.Fatalf("读取配置文件: %v", statErr)
	}

	lock.Lock()
	defer lock.Unlock()

	if _, err := toml.DecodeFile(envPath, &EnvSettings); err != nil {
		golog.Fatalf("加载配置文件：%v", err)
	}

	return &EnvSettings
}
