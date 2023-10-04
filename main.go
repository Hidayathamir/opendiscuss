package main

import (
	"os"

	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/router"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	db_connection := os.Getenv(constant.DB_CONNECTION)
	if db_connection == "" {
		panic("db connection not found in os environtment")
	}
	db, err := gorm.Open(mysql.Open(db_connection))
	if err != nil {
		panic(err)
	}

	r := router.GetRouter(db)
	if err := r.Run(); err != nil {
		panic(err)
	}
}
