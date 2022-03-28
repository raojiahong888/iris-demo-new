package datasource

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"iris-demo-new/config"
	"net/url"
	"os"
	"time"
)

var Instance *gorm.DB

func DB() *gorm.DB {
	if Instance != nil && Instance.DB().Ping() == nil {
		return Instance
	}
	return newConnect()
}

func newConnect() *gorm.DB {
	timezone := url.QueryEscape(config.Setting.App.Timezone)
	resource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s", config.Setting.Mysql.User, config.Setting.Mysql.Password, config.Setting.Mysql.Host, config.Setting.Mysql.Port, config.Setting.Mysql.Database, config.Setting.Mysql.Charset, timezone)
	db, err := gorm.Open("mysql", resource)
	if err != nil {
		fmt.Printf("mysql connect error:%s\r\n", err)
		os.Exit(-1)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(config.Setting.Mysql.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.Setting.Mysql.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Second * config.Setting.Mysql.MaxConnLifetime)
	if config.Setting.Mysql.SqlLog {
		db.LogMode(true)
	}
	Instance = db
	return Instance
}