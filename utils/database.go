package utils

import (
	"fmt"

	"github.com/Hidayathamir/opendiscuss/environtment"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDBConnection() (*gorm.DB, error) {
	gormConfig := gorm.Config{Logger: logger.Default.LogMode(logger.Info)}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		environtment.DB_USER, environtment.DB_PASSWORD, environtment.DB_HOST,
		environtment.DB_PORT, environtment.DB_NAME,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gormConfig)
	if err != nil {
		return nil, errors.Wrap(err, "error open db connection")
	}
	return db, nil
}
