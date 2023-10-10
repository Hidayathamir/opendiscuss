package controller

import (
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/gin-gonic/gin"
)

func (cc *CommentController) VoteThumbsUp(ctx *gin.Context) {
	cc.voteThumbs(ctx, constant.VoteOptionThumbsUp)
}
