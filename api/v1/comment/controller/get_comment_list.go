package controller

import (
	"net/http"

	"github.com/Hidayathamir/opendiscuss/api/v1/comment/dto"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (cc *CommentController) GetCommentList(ctx *gin.Context) {
	comments, err := cc.service.GetCommentList(ctx)
	if err != nil {
		err = errors.Wrap(err, "error get comment list")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	res := dto.ResGetCommentList{
		Comments: comments,
	}

	utils.WriteResponse(ctx, http.StatusOK, res, nil)
}
