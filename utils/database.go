package utils

import (
	"errors"
	"os"

	"github.com/Hidayathamir/opendiscuss/constant"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDBConnection() (*gorm.DB, error) {
	db_connection := os.Getenv(constant.DB_CONNECTION)
	if db_connection == "" {
		return nil, errors.New("db connection not found in os environtment")
	}
	gormConfig := gorm.Config{Logger: logger.Default.LogMode(logger.Info)}
	db, err := gorm.Open(mysql.Open(db_connection), &gormConfig)
	if err != nil {
		return nil, err
	}
	return db, nil
}
