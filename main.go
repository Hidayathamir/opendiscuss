package main

import (
	"os"

	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/router"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	if err := utils.SetTimeZone(); err != nil {
		panic(err)
	}
	db_connection := os.Getenv(constant.DB_CONNECTION)
	if db_connection == "" {
		panic("db connection not found in os environtment")
	}
	gormConfig := gorm.Config{Logger: logger.Default.LogMode(logger.Info)}
	db, err := gorm.Open(mysql.Open(db_connection), &gormConfig)
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	router.AddRouter(db, r)
	if err := r.Run(); err != nil {
		panic(err)
	}
}
