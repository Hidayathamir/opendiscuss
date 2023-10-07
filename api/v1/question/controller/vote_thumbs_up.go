package controller

import (
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/gin-gonic/gin"
)

func (qc *QuestionController) VoteThumbsUp(ctx *gin.Context) {
	qc.voteThumbs(ctx, constant.VoteOptionThumbsUp)
}
