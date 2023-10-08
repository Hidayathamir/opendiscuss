package controller

import (
	"net/http"
	"strconv"

	"github.com/Hidayathamir/opendiscuss/api/v1/answer/dto"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (ac *AnswerController) GetAnswerByID(ctx *gin.Context) {
	answerID, err := strconv.Atoi(ctx.Param("answerid"))
	if err != nil {
		err = errors.Wrap(err, "error convert answer id")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	answer, err := ac.service.GetAnswerByID(ctx, answerID)
	if err != nil {
		err = errors.Wrap(err, "error get answer")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	res := dto.ResGetAnswerByID{
		Answer: answer,
	}

	utils.WriteResponse(ctx, http.StatusOK, res, nil)
}
