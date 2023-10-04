package router

import (
	"net/http"

	questionrouter "github.com/Hidayathamir/opendiscuss/api/v1/question/router"
	"github.com/Hidayathamir/opendiscuss/utils"
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

	r.NoRoute(func(ctx *gin.Context) {
		utils.WriteResponse(ctx, http.StatusNotFound, nil, "page not found")
	})
}
