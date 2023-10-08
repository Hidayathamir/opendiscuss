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

func (qc *QuestionController) DeleteQuestionByID(ctx *gin.Context) {
	questionID, err := strconv.Atoi(ctx.Param("questionid"))
	if err != nil {
		err = errors.Wrap(err, "error convert question id")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	req := dto.ReqDeleteQuestionByID{
		UserID:     ctx.GetInt(constant.CTX_USER_ID),
		QuestionID: questionID,
	}

	questionID, err = qc.service.DeleteQuestionByID(ctx, req)
	if err != nil {
		err = errors.Wrap(err, "error delete question by id")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	res := dto.ResDeleteQuestionByID{
		QuestionID: questionID,
	}

	utils.WriteResponse(ctx, http.StatusOK, res, nil)
}
