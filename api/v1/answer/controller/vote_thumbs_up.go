package controller

import (
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/gin-gonic/gin"
)

func (ac *AnswerController) VoteThumbsUp(ctx *gin.Context) {
	ac.voteThumbs(ctx, constant.VoteOptionThumbsUp)
}
