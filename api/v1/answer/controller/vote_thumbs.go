package controller

import (
	"net/http"
	"strconv"

	"github.com/Hidayathamir/opendiscuss/api/v1/answer/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (ac *AnswerController) voteThumbs(ctx *gin.Context, voteOptionID constant.VoteOption) {
	isVoteOptionValid := voteOptionID == constant.VoteOptionThumbsUp ||
		voteOptionID == constant.VoteOptionThumbsDown

	if !isVoteOptionValid {
		err := errors.New("error vote option not valid")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	answerID, err := strconv.Atoi(ctx.Param("answerid"))
	if err != nil {
		err = errors.Wrap(err, "error convert answer id")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	req := dto.ReqVoteThumbs{
		UserID:   ctx.GetInt(constant.CTX_USER_ID),
		AnswerID: answerID,
	}

	if voteOptionID == constant.VoteOptionThumbsUp {
		answerID, err = ac.service.VoteThumbsUp(ctx, req)
		if err != nil {
			err = errors.Wrap(err, "error vote thumbs up")
			utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
			return
		}
	}

	if voteOptionID == constant.VoteOptionThumbsDown {
		answerID, err = ac.service.VoteThumbsDown(ctx, req)
		if err != nil {
			err = errors.Wrap(err, "error vote thumbs down")
			utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
			return
		}
	}

	res := dto.ResVoteThumbs{
		AnswerID: answerID,
	}

	utils.WriteResponse(ctx, http.StatusOK, res, nil)
}
