package controller

import (
	"net/http"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/gin-gonic/gin"
)

func (qc *QuestionController) CreateQuestion(ctx *gin.Context) {
	req := dto.ReqCreateQuestion{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, constant.INVALID_REQUEST_BODY)
		return
	}

	questionID, err := qc.service.CreateQuestion(ctx, req)
	if err != nil {
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	res := dto.ResCreateQuestion{
		QuestionID: questionID,
	}

	utils.WriteResponse(ctx, http.StatusOK, res, nil)
}
