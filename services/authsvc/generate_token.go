package authsvc

import (
	"net/http"

	"github.com/BearTS/go-gin-monolith/constants"
	"github.com/BearTS/go-gin-monolith/database"
	"github.com/BearTS/go-gin-monolith/dbops/gorm/admin"
	"github.com/BearTS/go-gin-monolith/dbops/gorm/users"
	"github.com/BearTS/go-gin-monolith/models"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (g *authSvcImpl) GenerateToken(c *gin.Context, req TokenReq) (utils.BaseResponse, TokenRes, error) {
	var baseRes utils.BaseResponse
	var res TokenRes
	var err error

	switch req.Type {

	case constants.TokenTypes.USER:
		{
			baseRes, res, err := g.userTokenGeneration(c, req)
			if err != nil {
				return baseRes, res, errors.Wrap(err, "[GenerateToken][userTokenGeneration]")
			}
			return baseRes, res, err
		}
	case constants.TokenTypes.ADMIN:
		{
			baseRes, res, err := g.adminTokenGeneration(c, req)
			if err != nil {
				return baseRes, res, errors.Wrap(err, "[GenerateToken][adminTokenGeneration]")
			}
			return baseRes, res, err
		}
	default:
		{
			baseRes.Success = false
			baseRes.Message = "invalid token type"
			baseRes.StatusCode = http.StatusUnprocessableEntity
			return baseRes, res, err
		}
	}
}

func (g *authSvcImpl) userTokenGeneration(c *gin.Context, req TokenReq) (utils.BaseResponse, TokenRes, error) {
	var baseRes utils.BaseResponse
	var res TokenRes
	var err error
	var authData models.AuthData

	authData.Type = req.Type

	//get sandbox
	gormDB, _ := database.Connection()
	userGorm := users.Gorm(gormDB)
	userData, err := userGorm.GetUserDetailsByPID(c, req.UserID)
	if err != nil {
		return baseRes, res, errors.Wrap(err, "[onboardingTokenGeneration][GetUserDetailsByPID]")
	}
	authData.Sandbox = userData.IsSandbox
	authData.UserPID = userData.PID

	tokenRes, err := g.CreateToken(authData)
	if err != nil {
		return baseRes, res, errors.Wrap(err, "[onboardingTokenGeneration][CreateToken][onboarding]")
	}
	res.AccesssToken = tokenRes.AccessToken
	res.RefreshToken = tokenRes.RefreshToken
	res.AccessTokenExp = tokenRes.AtExpires
	res.RefreshTokenExp = tokenRes.RtExpires
	res.AccesssTokenPID = tokenRes.TokenUuid
	res.RefreshTokenPID = tokenRes.RefreshUuid
	res.IsPhoneAvailable = true
	res.Type = req.Type
	res.UserID = userData.PID

	baseRes.Success = true
	baseRes.Message = "token generated successfully"
	baseRes.StatusCode = http.StatusOK

	return baseRes, res, err
}

func (g *authSvcImpl) adminTokenGeneration(c *gin.Context, req TokenReq) (utils.BaseResponse, TokenRes, error) {
	var baseRes utils.BaseResponse
	var res TokenRes
	var err error
	var authData models.AuthData

	authData.Type = req.Type

	//get sandbox
	gormDB, _ := database.Connection()
	adminGorm := admin.Gorm(gormDB)

	adminData, err := adminGorm.GetAdminByPID(c, req.UserID)
	if err != nil {
		return baseRes, res, errors.Wrap(err, "[onboardingTokenGeneration][GetUserDetailsByPID]")
	}
	authData.Sandbox = adminData.IsSandbox
	authData.AdminPID = adminData.PID

	tokenRes, err := g.CreateToken(authData)
	if err != nil {
		return baseRes, res, errors.Wrap(err, "[onboardingTokenGeneration][CreateToken][onboarding]")
	}
	res.AccesssToken = tokenRes.AccessToken
	res.RefreshToken = tokenRes.RefreshToken
	res.AccessTokenExp = tokenRes.AtExpires
	res.RefreshTokenExp = tokenRes.RtExpires
	res.AccesssTokenPID = tokenRes.TokenUuid
	res.RefreshTokenPID = tokenRes.RefreshUuid
	res.IsPhoneAvailable = true
	res.Type = req.Type
	res.AdminID = adminData.PID

	baseRes.Success = true
	baseRes.Message = "token generated successfully"
	baseRes.StatusCode = http.StatusOK

	return baseRes, res, err
}
