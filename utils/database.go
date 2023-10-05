package utils

import (
	"github.com/Hidayathamir/opendiscuss/environtment"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDBConnection() (*gorm.DB, error) {
	gormConfig := gorm.Config{Logger: logger.Default.LogMode(logger.Info)}
	db, err := gorm.Open(mysql.Open(environtment.DB_CONNECTION), &gormConfig)
	if err != nil {
		return nil, err
	}
	return db, nil
}
