package db

import (
	"fmt"
	"github.com/kataras/golog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"im-project/config"
	"sync"
	"time"
)

var db *gorm.DB
var once sync.Once

func NewConnection() *gorm.DB {
	settings := config.GetEnv().Mysql

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True", settings.DbUser,
		settings.DbPassword, settings.DbHost, settings.DbPort, settings.DbDatabase, settings.DbCharset)
	golog.Debugf("dsn: %s", dsn)

	conn, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		golog.Fatalf("连接数据库失败: %v", err)
		return nil
	}

	sqlDB, err := conn.DB()
	sqlDB.SetMaxIdleConns(settings.PoolMaxIdleConns)
	sqlDB.SetMaxOpenConns(settings.PoolMaxConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return conn
}

func MysqlInit() *gorm.DB {
	once.Do(func() {
		db = NewConnection()
		golog.Debugf("### DB 已初始化")
	})
	return db
}

func GetDB() *gorm.DB {
	return MysqlInit()
}
