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

func (qc *QuestionController) UpdateQuestionByID(ctx *gin.Context) {
	questionID, err := strconv.Atoi(ctx.Param("questionid"))
	if err != nil {
		err = errors.Wrap(err, "error convert question id")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	req := dto.ReqUpdateQuestionByID{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		err = errors.Wrap(err, constant.ERR_INVALID_REQUEST_BODY)
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	req.QuestionID = questionID
	req.UserID = ctx.GetInt(constant.CTX_USER_ID)

	questionID, err = qc.service.UpdateQuestionByID(ctx, req)
	if err != nil {
		err = errors.Wrap(err, "error update question by id")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	res := dto.ResUpdateQuestionByID{
		QuestionID: questionID,
	}

	utils.WriteResponse(ctx, http.StatusOK, res, nil)
}
