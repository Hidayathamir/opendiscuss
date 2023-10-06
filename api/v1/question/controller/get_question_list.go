package controller

import (
	"net/http"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (qc *QuestionController) GetQuestionList(ctx *gin.Context) {
	questions, err := qc.service.GetQuestionList(ctx)
	if err != nil {
		err = errors.Wrap(err, "error get question list")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	res := dto.ResGetQuestionList{
		Questions: questions,
	}

	utils.WriteResponse(ctx, http.StatusOK, res, nil)
}
