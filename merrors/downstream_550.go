package merrors

import (
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*                              DOWNSTREAM ERROR                              */
/* -------------------------------------------------------------------------- */
func Downstream(ctx *gin.Context, err string) {
	var res utils.BaseResponse
	var smerror Error
	errorCode := 550
	smerror.Code = errorCode
	smerror.Type = errorType.Downstream
	smerror.Message = err
	res.Error = smerror
	ctx.JSON(errorCode, res)
	ctx.Abort()
}
