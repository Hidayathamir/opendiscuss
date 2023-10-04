package router

import (
	questionrouter "github.com/Hidayathamir/opendiscuss/api/v1/question/router"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddRouter(db *gorm.DB, r *gin.Engine) {
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			questionrouter.AddQuestionRouter(db, v1)
		}
	}
}
