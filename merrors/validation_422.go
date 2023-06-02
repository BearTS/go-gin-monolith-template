package merrors

import (
	"net/http"

	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
)

type Error struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

/* -------------------------------------------------------------------------- */
/*                            VALIDATION ERROR 422                            */
/* -------------------------------------------------------------------------- */
func Validation(ctx *gin.Context, err string) {
	var res utils.BaseResponse
	var smerror Error
	errorCode := http.StatusUnprocessableEntity

	smerror.Code = errorCode
	smerror.Type = errorType.validation
	smerror.Message = err

	res.Error = smerror

	ctx.JSON(errorCode, res)
	ctx.Abort()
}
