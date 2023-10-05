package controller

import (
	"net/http"

	"github.com/Hidayathamir/opendiscuss/api/v1/user/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (uc *UserController) RegisterUser(ctx *gin.Context) {
	req := dto.ReqRegisterUser{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		err = errors.Wrap(err, constant.ERR_INVALID_REQUEST_BODY)
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	userID, err := uc.service.RegisterUser(ctx, req)
	if err != nil {
		err = errors.Wrap(err, "error register user")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	res := dto.ResRegisterUser{
		UserID: userID,
	}

	utils.WriteResponse(ctx, http.StatusOK, res, nil)
}
