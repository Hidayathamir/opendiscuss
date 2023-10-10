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

func (cc *CommentController) DeleteCommentByID(ctx *gin.Context) {
	commentID, err := strconv.Atoi(ctx.Param("commentid"))
	if err != nil {
		err = errors.Wrap(err, "error convert comment id")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	req := dto.ReqDeleteCommentByID{
		UserID:    ctx.GetInt(constant.CTX_USER_ID),
		CommentID: commentID,
	}

	commentID, err = cc.service.DeleteCommentByID(ctx, req)
	if err != nil {
		err = errors.Wrap(err, "error delete comment by id")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	res := dto.ResDeleteCommentByID{
		CommentID: commentID,
	}

	utils.WriteResponse(ctx, http.StatusOK, res, nil)
}
