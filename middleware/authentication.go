package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func Authenticate(ctx *gin.Context) {
	auth := ctx.Request.Header.Get(constant.HEADER_AUTHORIZATION)
	if len(auth) == 0 {
		err := errors.New("error auth token not found")
		err = errors.Wrap(err, "please add header 'Authorization: Bearer your-token'")
		utils.WriteResponse(ctx, http.StatusUnauthorized, nil, err)
		ctx.Abort()
		return
	}

	authSlice := strings.Split(auth, " ")

	if len(authSlice) != 2 {
		err := errors.New("error token format")
		err = errors.Wrap(err, "token format should be 'Bearer your-token'")
		utils.WriteResponse(ctx, http.StatusUnauthorized, nil, err)
		ctx.Abort()
		return
	}

	isTokenTypeBearer := authSlice[0] != "Bearer"
	if isTokenTypeBearer {
		err := errors.New("error authentication")
		err = errors.Wrap(err, "token type should be Bearer")
		utils.WriteResponse(ctx, http.StatusUnauthorized, nil, err)
		ctx.Abort()
		return
	}

	claims, err := utils.ValidateUserJWTToken(authSlice[1])
	if err != nil {
		err = errors.Wrap(err, "error validate user jwt token")
		utils.WriteResponse(ctx, http.StatusUnauthorized, nil, err)
		ctx.Abort()
		return
	}

	userIDAny, ok := claims[constant.JWT_CLAIM_USER_ID]
	if !ok {
		err = errors.New("error user_id not found in token")
		utils.WriteResponse(ctx, http.StatusUnauthorized, nil, err)
		ctx.Abort()
		return
	}

	userIDStr, ok := userIDAny.(string)
	if !ok {
		err = errors.New("error user_id in token is not string")
		utils.WriteResponse(ctx, http.StatusUnauthorized, nil, err)
		ctx.Abort()
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		err = errors.New("error convert user_id to int")
		utils.WriteResponse(ctx, http.StatusUnauthorized, nil, err)
		ctx.Abort()
		return
	}

	ctx.Set(constant.CTX_USER_ID, userID)
}
