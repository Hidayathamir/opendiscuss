package controller

import (
	"net/http"
	"strconv"

	"github.com/Hidayathamir/opendiscuss/api/v1/comment/dto"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (cc *CommentController) GetCommentListByAnswerID(ctx *gin.Context) {
	answerID, err := strconv.Atoi(ctx.Param("answerid"))
	if err != nil {
		err = errors.Wrap(err, "error convert question id")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	comments, err := cc.service.GetCommentListByAnswerID(ctx, answerID)
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

func (cc *CommentController) GetSubCommentListByCommentID(ctx *gin.Context) {
	commentID, err := strconv.Atoi(ctx.Param("commentid"))
	if err != nil {
		err = errors.Wrap(err, "error convert comment id")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	subComments, err := cc.service.GetSubCommentListByCommentID(ctx, commentID)
	if err != nil {
		err = errors.Wrap(err, "error get sub comment list")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	res := dto.ResGetCommentList{
		Comments: subComments,
	}

	utils.WriteResponse(ctx, http.StatusOK, res, nil)
}
