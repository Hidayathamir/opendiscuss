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

func (ac *AnswerController) DeleteAnswerByID(ctx *gin.Context) {
	answerID, err := strconv.Atoi(ctx.Param("answerid"))
	if err != nil {
		err = errors.Wrap(err, "error convert answer id")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	req := dto.ReqDeleteAnswerByID{
		UserID:   ctx.GetInt(constant.CTX_USER_ID),
		AnswerID: answerID,
	}

	answerID, err = ac.service.DeleteAnswerByID(ctx, req)
	if err != nil {
		err = errors.Wrap(err, "error delete answer by id")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	res := dto.ResDeleteAnswerByID{
		AnswerID: answerID,
	}

	utils.WriteResponse(ctx, http.StatusOK, res, nil)
}
