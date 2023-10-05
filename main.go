package main

import (
	"github.com/Hidayathamir/opendiscuss/environtment"
	"github.com/Hidayathamir/opendiscuss/router"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	if err := environtment.InitEnv(); err != nil {
		panic(err)
	}

	if err := utils.SetTimeZone(); err != nil {
		panic(err)
	}

	db, err := utils.GetDBConnection()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	router.AddRouter(db, r)
	if err := r.Run(); err != nil {
		panic(err)
	}
}
