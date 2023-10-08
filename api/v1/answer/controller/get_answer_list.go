package controller

import (
	"net/http"
	"strconv"

	"github.com/Hidayathamir/opendiscuss/api/v1/answer/dto"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (ac *AnswerController) GetAnswerListByQuestionID(ctx *gin.Context) {
	questionID, err := strconv.Atoi(ctx.Param("questionid"))
	if err != nil {
		err = errors.Wrap(err, "error convert question id")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	answers, err := ac.service.GetAnswerListByQuestionID(ctx, questionID)
	if err != nil {
		err = errors.Wrap(err, "error get answer list by question id")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	res := dto.ResGetAnswerListByQuestionID{
		Answers: answers,
	}

	utils.WriteResponse(ctx, http.StatusOK, res, nil)
}
