package utils

import "github.com/gin-gonic/gin"

func WriteResponse(ctx *gin.Context, code int, data interface{}, _error interface{}) {
	ctx.JSON(code, gin.H{
		"data":  data,
		"error": _error,
	})
}
