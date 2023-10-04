package main

import (
	"github.com/Hidayathamir/opendiscuss/router"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open(""))
	if err != nil {
		panic(err)
	}
	r := router.GetRouter(db)
	if err := r.Run(); err != nil {
		panic(err)
	}
}
