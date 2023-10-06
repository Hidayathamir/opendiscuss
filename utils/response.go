package utils

import "github.com/gin-gonic/gin"

func WriteResponse(ctx *gin.Context, code int, data interface{}, _error error) {
	body := gin.H{
		"data":  data,
		"error": nil,
	}
	if _error != nil {
		body["error"] = _error.Error()
	}
	ctx.JSON(code, body)
}
