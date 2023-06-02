package merrors

import (
	"net/http"

	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
)

func HandleServiceCodes(ctx *gin.Context, baseRes utils.BaseResponse) {
	switch baseRes.StatusCode {
	case http.StatusUnauthorized:
		{
			Unauthorized(ctx, baseRes.Message)
		}
	case http.StatusForbidden:
		{
			Forbidden(ctx, baseRes.Message)
		}
	case http.StatusServiceUnavailable:
		{
			ServiceUnavailable(ctx, baseRes.Message)
		}
	case http.StatusConflict:
		{
			Conflict(ctx, baseRes.Message)
		}
	case http.StatusUnprocessableEntity:
		{
			Validation(ctx, baseRes.Message)
		}
	case 550:
		{
			Downstream(ctx, baseRes.Message)
		}

	default:
		{
			InternalServer(ctx, baseRes.Message)
		}
	}
}
