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

func (cc *CommentController) UpdateCommentByID(ctx *gin.Context) {
	commentID, err := strconv.Atoi(ctx.Param("commentid"))
	if err != nil {
		err = errors.Wrap(err, "error convert comment id")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	req := dto.ReqUpdateCommentByID{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		err = errors.Wrap(err, constant.ERR_INVALID_REQUEST_BODY)
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	req.CommentID = commentID
	req.UserID = ctx.GetInt(constant.CTX_USER_ID)

	commentID, err = cc.service.UpdateCommentByID(ctx, req)
	if err != nil {
		err = errors.Wrap(err, "error update comment by id")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	res := dto.ResUpdateCommentByID{
		CommentID: commentID,
	}

	utils.WriteResponse(ctx, http.StatusOK, res, nil)
}
