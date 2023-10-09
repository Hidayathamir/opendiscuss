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

func (ac *AnswerController) UpdateAnswerByID(ctx *gin.Context) {
	answerID, err := strconv.Atoi(ctx.Param("answerid"))
	if err != nil {
		err = errors.Wrap(err, "error convert answer id")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	req := dto.ReqUpdateAnswerByID{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		err = errors.Wrap(err, constant.ERR_INVALID_REQUEST_BODY)
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	req.AnswerID = answerID
	req.UserID = ctx.GetInt(constant.CTX_USER_ID)

	answerID, err = ac.service.UpdateAnswerByID(ctx, req)
	if err != nil {
		err = errors.Wrap(err, "error update answer by id")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	res := dto.ResUpdateAnswerByID{
		AnswerID: answerID,
	}

	utils.WriteResponse(ctx, http.StatusOK, res, nil)
}
