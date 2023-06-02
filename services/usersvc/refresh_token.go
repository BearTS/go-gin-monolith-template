package usersvc

import (
	"errors"
	"log"
	"net/http"

	"github.com/BearTS/go-gin-monolith/constants"
	"github.com/BearTS/go-gin-monolith/database/tables"
	"github.com/BearTS/go-gin-monolith/services/authsvc"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (g *UserSvcImpl) RefreshToken(c *gin.Context) (utils.BaseResponse, tables.Users, error) {
	var baseRes utils.BaseResponse
	var res tables.Users
	var err error

	refreshTokeData, err := utils.GetRefreshTokenData(c)
	if err != nil {
		return baseRes, res, err
	}

	if refreshTokeData.UserPID == "" {
		baseRes.Success = false
		baseRes.Message = "invalid token"
		baseRes.StatusCode = http.StatusUnauthorized

		return baseRes, res, err
	}

	// Generate a new token pair.
	// Get Details from Db
	user, err := g.usersGorm.GetUserDetailsByPID(c, refreshTokeData.UserPID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("[Critical][Should Not Happen] User not found") // Should not happen
			baseRes.StatusCode = http.StatusNotFound
			baseRes.Message = constants.Errors.VerifyOTP.InvalidOTP
			return baseRes, res, err
		}
		return baseRes, res, err
	}

	var authData authsvc.TokenReq
	authData.UserID = user.PID

	// if user.MobileNumber == "" {
	// 	authData.Type = constants.TokenTypes.ONBOARDING
	// 	_, token, err := g.authSvc.GenerateToken(c, authData)
	// 	if err != nil {
	// 		return baseRes, res, err
	// 	}

	// 	// Add Success Response
	// 	baseRes.Success = true
	// 	baseRes.StatusCode = http.StatusOK
	// 	baseRes.Message = "User Created Successfully"
	// 	baseRes.Data = token

	// 	return baseRes, res, err
	// }

	authData.Type = constants.TokenTypes.USER
	_, token, err := g.authSvc.GenerateToken(c, authData)
	if err != nil {
		return baseRes, res, err
	}

	// Add Success Response
	baseRes.Success = true
	baseRes.StatusCode = http.StatusOK
	baseRes.Message = "User Created Successfully"
	baseRes.Data = token

	return baseRes, res, err
}
