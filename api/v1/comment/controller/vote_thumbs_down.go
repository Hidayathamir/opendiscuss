package controller

import (
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/gin-gonic/gin"
)

func (cc *CommentController) VoteThumbsDown(ctx *gin.Context) {
	cc.voteThumbs(ctx, constant.VoteOptionThumbsDown)
}
