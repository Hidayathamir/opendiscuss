package controller

import (
	"net/http"

	"github.com/Hidayathamir/opendiscuss/api/v1/user/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/gin-gonic/gin"
)

func (uc *UserController) LoginUser(ctx *gin.Context) {
	req := dto.ReqLoginUser{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, constant.INVALID_REQUEST_BODY)
		return
	}

	token, err := uc.service.LoginUser(ctx, req)
	if err != nil {
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	res := dto.ResLoginUser{
		Token: token,
	}

	utils.WriteResponse(ctx, http.StatusOK, res, nil)
}
