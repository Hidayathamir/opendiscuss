package controller

import (
	"net/http"
	"strconv"

	"github.com/Hidayathamir/opendiscuss/api/v1/comment/dto"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (cc *CommentController) GetCommentByID(ctx *gin.Context) {
	commentID, err := strconv.Atoi(ctx.Param("commentid"))
	if err != nil {
		err = errors.Wrap(err, "error convert comment id")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	comment, err := cc.service.GetCommentByID(ctx, commentID)
	if err != nil {
		err = errors.Wrap(err, "error get comment")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	res := dto.ResGetComment{
		Comment: comment,
	}

	utils.WriteResponse(ctx, http.StatusOK, res, nil)
}
