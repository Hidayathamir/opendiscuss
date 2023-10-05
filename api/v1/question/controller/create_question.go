package controller

import (
	"net/http"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (qc *QuestionController) CreateQuestion(ctx *gin.Context) {
	req := dto.ReqCreateQuestion{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		err = errors.Wrap(err, constant.ERR_INVALID_REQUEST_BODY)
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	req.UserID = ctx.GetInt("user_id")

	questionID, err := qc.service.CreateQuestion(ctx, req)
	if err != nil {
		err = errors.Wrap(err, "error create question")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	res := dto.ResCreateQuestion{
		QuestionID: questionID,
	}

	utils.WriteResponse(ctx, http.StatusOK, res, nil)
}
