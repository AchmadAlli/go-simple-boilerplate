package database

import (
	"fmt"

	"github.com/achmadAlli/go-simple-boilerplate/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlInstance(env config.MySQLEnv) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		env.Username,
		env.Password,
		env.Host,
		env.Port,
		env.Schema,
	)

	instance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(fmt.Sprintf("unable to open connection to database: %v", err))
	}

	sqlDB, err := instance.DB()
	if err != nil {
		panic(fmt.Sprintf("unable get database instance: %v", err))
	}

	if err := sqlDB.Ping(); err != nil {
		panic(fmt.Sprintf("failed to ping database: %v", err))
	}

	if env.Debug {
		instance = instance.Debug()
	}

	return instance
}
