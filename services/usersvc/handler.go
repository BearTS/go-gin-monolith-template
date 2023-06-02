package usersvc

import (
	"github.com/BearTS/go-gin-monolith/database/tables"
	"github.com/BearTS/go-gin-monolith/dbops/gorm/otp_verifications"
	"github.com/BearTS/go-gin-monolith/dbops/gorm/users"
	"github.com/BearTS/go-gin-monolith/services/authsvc"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
)

type UserSvcImpl struct {
	usersGorm           users.GormInterface
	otpVerificationGorm otp_verifications.GormInterface
	authSvc             authsvc.Interface
}

// interface
type Interface interface {
	SendOTP(c *gin.Context, req SendOTPReq) (utils.BaseResponse, tables.Users, error)
	VerifyOTP(c *gin.Context, req VerifyOTPReq) (utils.BaseResponse, tables.Users, error)
	ResendOTP(c *gin.Context, req ResendOTPReq) (utils.BaseResponse, tables.Users, error)
}

func Handler(userGorm users.GormInterface, otpVerificationGorm otp_verifications.GormInterface, authSvc authsvc.Interface) Interface {
	return &UserSvcImpl{
		usersGorm:           userGorm,
		otpVerificationGorm: otpVerificationGorm,
		authSvc:             authSvc,
	}
}
