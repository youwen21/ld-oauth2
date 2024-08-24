package conf

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type MysqlCfg struct {
	HOST     string `toml:"host" mapstructure:"host"`
	USERNAME string `toml:"username" mapstructure:"username"`
	PASSWORD string `toml:"password" mapstructure:"password"`
	PORT     string `toml:"port" mapstructure:"port"`
	DATABASE string `toml:"database" mapstructure:"database"`
	CHARSET  string `toml:"charset" mapstructure:"charset"`

	instance *gorm.DB
}

func (cfg *MysqlCfg) initConnPool() {
	dsn := cfg.GetDsn()
	mysqlDb := mysql.Open(dsn)

	db, err := gorm.Open(mysqlDb, &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDb, err := db.DB()
	if err != nil {
		panic("failed to connect database")
	}
	sqlDb.SetMaxIdleConns(5)
	sqlDb.SetMaxOpenConns(100)
	cfg.instance = db
}

func (cfg *MysqlCfg) GetDb() *gorm.DB {
	return cfg.instance
}

func (cfg *MysqlCfg) GetSession() *gorm.DB {
	if os.Getenv("APP_ENV") != "prod" {
		return cfg.instance.Debug()
	}

	return cfg.instance.Session(&gorm.Session{})
}

func (cfg *MysqlCfg) GetDsn() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=%s&loc=Local&parseTime=true", cfg.USERNAME, cfg.PASSWORD, cfg.HOST, cfg.PORT, cfg.DATABASE, cfg.CHARSET)
	return dsn
}
