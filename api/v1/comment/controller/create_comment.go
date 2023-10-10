package controller

import (
	"net/http"
	"strconv"

	"github.com/Hidayathamir/opendiscuss/api/v1/comment/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (cc *CommentController) CreateComment(ctx *gin.Context) {
	req := dto.ReqCreateComment{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		err = errors.Wrap(err, constant.ERR_INVALID_REQUEST_BODY)
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	req.UserID = ctx.GetInt(constant.CTX_USER_ID)

	answerID, err := strconv.Atoi(ctx.Param("answerid"))
	if err != nil {
		err = errors.Wrap(err, "error convert question id")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	req.AnswerID = answerID

	commentID, err := cc.service.CreateComment(ctx, req)
	if err != nil {
		err = errors.Wrap(err, "error create comment")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	res := dto.ResCreateComment{
		CommentID: commentID,
	}

	utils.WriteResponse(ctx, http.StatusOK, res, nil)
}
