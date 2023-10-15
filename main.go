package main

import (
	"log"
	"time"

	"github.com/Hidayathamir/opendiscuss/environtment"
	"github.com/Hidayathamir/opendiscuss/router"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(errors.Wrap(err, "error load .env file"))
	}

	if err := environtment.InitEnv(); err != nil {
		panic(errors.Wrap(err, "error init environtment"))
	}

	var db *gorm.DB
	var err error
	for {
		db, err = utils.GetDBConnection()
		if err == nil {
			break
		}

		err = errors.Wrap(err, "error get db connection, waiting 5s before connect again")
		log.Println(err)

		time.Sleep(5 * time.Second)
	}

	r := gin.Default()
	r.Use(cors.Default())
	router.AddRouter(db, r)
	if err := r.Run(); err != nil {
		panic(errors.Wrap(err, "error gin router run"))
	}
}
