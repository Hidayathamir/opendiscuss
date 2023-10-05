package main

import (
	"github.com/Hidayathamir/opendiscuss/environtment"
	"github.com/Hidayathamir/opendiscuss/router"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(errors.Wrap(err, "error load .env file"))
	}

	if err := environtment.InitEnv(); err != nil {
		panic(errors.Wrap(err, "error init environtment"))
	}

	if err := utils.SetTimeZone(); err != nil {
		panic(errors.Wrap(err, "error set time zone"))
	}

	db, err := utils.GetDBConnection()
	if err != nil {
		panic(errors.Wrap(err, "error get db connection"))
	}

	r := gin.Default()
	router.AddRouter(db, r)
	if err := r.Run(); err != nil {
		panic(errors.Wrap(err, "error gin router run"))
	}
}
