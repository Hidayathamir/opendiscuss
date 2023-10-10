package router

import (
	"errors"
	"net/http"

	answerrouter "github.com/Hidayathamir/opendiscuss/api/v1/answer/router"
	commentrouter "github.com/Hidayathamir/opendiscuss/api/v1/comment/router"
	questionrouter "github.com/Hidayathamir/opendiscuss/api/v1/question/router"
	userrouter "github.com/Hidayathamir/opendiscuss/api/v1/user/router"
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
			userrouter.AddUserRouter(db, v1)
			answerrouter.AddAnswerRouter(db, v1)
			commentrouter.AddCommentRouter(db, v1)
		}
	}

	r.NoRoute(func(ctx *gin.Context) {
		utils.WriteResponse(ctx, http.StatusNotFound, nil, errors.New("page not found"))
	})
}
