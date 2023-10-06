package controller

import (
	"net/http"

	"github.com/Hidayathamir/opendiscuss/api/v1/user/dto"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (uc *UserController) LoginUser(ctx *gin.Context) {
	req := dto.ReqLoginUser{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		err = errors.Wrap(err, constant.ERR_INVALID_REQUEST_BODY)
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	token, err := uc.service.LoginUser(ctx, req)
	if err != nil {
		err = errors.Wrap(err, "error login user")
		utils.WriteResponse(ctx, http.StatusBadRequest, nil, err)
		return
	}

	res := dto.ResLoginUser{
		Token: token,
	}

	utils.WriteResponse(ctx, http.StatusOK, res, nil)
}
