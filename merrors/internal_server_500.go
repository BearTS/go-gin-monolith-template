package merrors

import (
	"net/http"

	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*                            INTERNAL SERVER ERROR                           */
/* -------------------------------------------------------------------------- */
func InternalServer(ctx *gin.Context, err string) {
	var res utils.BaseResponse
	var smerror Error
	errorCode := http.StatusInternalServerError

	smerror.Code = errorCode
	smerror.Type = errorType.server
	smerror.Message = err

	res.Error = smerror

	ctx.JSON(errorCode, res)
	ctx.Abort()
}
