package utils

import (
	"github.com/Hidayathamir/opendiscuss/environtment"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDBConnection() (*gorm.DB, error) {
	gormConfig := gorm.Config{Logger: logger.Default.LogMode(logger.Info)}
	db, err := gorm.Open(mysql.Open(environtment.DB_CONNECTION), &gormConfig)
	if err != nil {
		return nil, errors.Wrap(err, "error open db connection")
	}
	return db, nil
}
