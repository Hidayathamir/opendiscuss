package controller

import (
	"net/http"
	"strconv"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (qc *QuestionController) voteThumbs(ctx *gin.Context, voteOptionID constant.VoteOption) {
	isVoteOptionValid := voteOptionID == constant.VoteOptionThumbsUp ||
		voteOptionID == constant.VoteOptionThumbsDown

	if !isVoteOptionValid {
		err := errors.New("error vote option not valid")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	questionID, err := strconv.Atoi(ctx.Param("questionid"))
	if err != nil {
		err = errors.Wrap(err, "error convert question id")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	req := dto.ReqVoteThumbs{
		UserID:     ctx.GetInt(constant.CTX_USER_ID),
		QuestionID: questionID,
	}

	if voteOptionID == constant.VoteOptionThumbsUp {
		questionID, err = qc.service.VoteThumbsUp(ctx, req)
		if err != nil {
			err = errors.Wrap(err, "error vote thumbs up")
			utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
			return
		}
	}

	if voteOptionID == constant.VoteOptionThumbsDown {
		questionID, err = qc.service.VoteThumbsDown(ctx, req)
		if err != nil {
			err = errors.Wrap(err, "error vote thumbs down")
			utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
			return
		}
	}

	res := dto.ResVoteThumbs{
		QuestionID: questionID,
	}

	utils.WriteResponse(ctx, http.StatusOK, res, nil)
}
