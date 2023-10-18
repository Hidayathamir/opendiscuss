package middleware

import (
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, constant.ALLOWED_HEADER...)
	return cors.New(config)
}
