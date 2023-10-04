package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	return r
}
