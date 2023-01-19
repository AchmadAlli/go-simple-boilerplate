package config

import (
	"gorm.io/gorm"
)

type database struct {
	userDB *gorm.DB
}

var DBConfig *database

func LoadDBConfig() {
	if DBConfig == nil {
		DBConfig = &database{}
	}
}

func (db *database) SetUserDB(instance *gorm.DB) {
	db.userDB = instance
}

func (db *database) GetUserDB() *gorm.DB {
	return db.userDB
}
