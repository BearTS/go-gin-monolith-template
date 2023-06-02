package merrors

import (
	"net/http"

	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
)

func ServiceUnavailable(ctx *gin.Context, err string) {
	var res utils.BaseResponse
	var smerror Error
	errorCode := http.StatusServiceUnavailable
	smerror.Code = errorCode
	smerror.Type = errorType.ServiceUnavailable
	smerror.Message = err

	res.Error = smerror

	ctx.JSON(errorCode, res)
	ctx.Abort()
}
