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

func (cc *CommentController) voteThumbs(ctx *gin.Context, voteOptionID constant.VoteOption) {
	isVoteOptionValid := voteOptionID == constant.VoteOptionThumbsUp ||
		voteOptionID == constant.VoteOptionThumbsDown

	if !isVoteOptionValid {
		err := errors.New("error vote option not valid")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	commentID, err := strconv.Atoi(ctx.Param("commentid"))
	if err != nil {
		err = errors.Wrap(err, "error convert comment id")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	req := dto.ReqVoteThumbs{
		UserID:    ctx.GetInt(constant.CTX_USER_ID),
		CommentID: commentID,
	}

	if voteOptionID == constant.VoteOptionThumbsUp {
		commentID, err = cc.service.VoteThumbsUp(ctx, req)
		if err != nil {
			err = errors.Wrap(err, "error vote thumbs up")
			utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
			return
		}
	}

	if voteOptionID == constant.VoteOptionThumbsDown {
		commentID, err = cc.service.VoteThumbsDown(ctx, req)
		if err != nil {
			err = errors.Wrap(err, "error vote thumbs down")
			utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
			return
		}
	}

	res := dto.ResVoteThumbs{
		CommentID: commentID,
	}

	utils.WriteResponse(ctx, http.StatusOK, res, nil)
}
