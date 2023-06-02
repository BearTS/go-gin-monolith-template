package user

import (
	"net/http"

	"github.com/BearTS/go-gin-monolith/database/tables"
	"github.com/BearTS/go-gin-monolith/services/usersvc"
	"github.com/BearTS/go-gin-monolith/utils"
)

func sendOtpTransformer(data tables.Users) utils.BaseResponse {
	var res utils.BaseResponse
	var dataRes usersvc.SendOTPRes

	res.Success = true
	res.StatusCode = http.StatusOK
	res.Message = "OTP sent successfully"

	dataRes.Email = data.Email
	dataRes.UserPID = data.PID

	res.Data = dataRes

	return res
}
