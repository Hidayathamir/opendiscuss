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

func (ac *AnswerController) CreateAnswer(ctx *gin.Context) {
	req := dto.ReqCreateAnswer{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		err = errors.Wrap(err, constant.ERR_INVALID_REQUEST_BODY)
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	req.UserID = ctx.GetInt(constant.CTX_USER_ID)

	questionID, err := strconv.Atoi(ctx.Param("questionid"))
	if err != nil {
		err = errors.Wrap(err, "error convert question id")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	req.QuestionID = questionID

	answerID, err := ac.service.CreateAnswer(ctx, req)
	if err != nil {
		err = errors.Wrap(err, "error create answer")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	res := dto.ResCreateAnswer{
		AnswerID: answerID,
	}

	utils.WriteResponse(ctx, http.StatusOK, res, nil)
}
