package controller

import (
	"net/http"
	"strconv"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (qc *QuestionController) GetQuestionByID(ctx *gin.Context) {
	questionID, err := strconv.Atoi(ctx.Param("questionid"))
	if err != nil {
		err = errors.Wrap(err, "error convert question id")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	question, err := qc.service.GetQuestionByID(ctx, questionID)
	if err != nil {
		err = errors.Wrap(err, "error get question")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	res := dto.ResGetQuestion{
		Question: question,
	}

	utils.WriteResponse(ctx, http.StatusOK, res, nil)
}
