package user

import (
	"github.com/BearTS/go-gin-monolith/constants"
	"github.com/BearTS/go-gin-monolith/database"
	"github.com/BearTS/go-gin-monolith/dbops/gorm/otp_verifications"
	"github.com/BearTS/go-gin-monolith/dbops/gorm/users"
	"github.com/BearTS/go-gin-monolith/services/usersvc"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func validateSendOTPReq(c *gin.Context) (usersvc.SendOTPReq, error) {
	var req usersvc.SendOTPReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return req, err
	}
	return req, err
}

func validateResendOTPReq(c *gin.Context) (usersvc.ResendOTPReq, error) {
	var req usersvc.ResendOTPReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return req, err
	}
	return req, err
}

func validateVerifyOTPReq(c *gin.Context) (usersvc.VerifyOTPReq, error) {
	var req usersvc.VerifyOTPReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return req, err
	}

	if err := verifyUserInDb(c, req.UserPID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return req, errors.New(constants.Errors.VerifyOTP.ResendOTP)
		}
		return req, err
	}

	if err := verifyOtpInDb(c, req.UserPID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return req, errors.New(constants.Errors.VerifyOTP.ResendOTP)
		}
		return req, err
	}

	return req, err
}

func verifyUserInDb(c *gin.Context, pid string) error {
	var err error

	gormDB, _ := database.Connection()
	usersGorm := users.Gorm(gormDB)

	_, err = usersGorm.GetUserDetailsByPID(c, pid)

	if err != nil {
		return err
	}

	return err
}

func verifyOtpInDb(c *gin.Context, pid string) error {
	var err error

	gormDB, _ := database.Connection()
	otpVerificationGorm := otp_verifications.Gorm(gormDB)

	_, err = otpVerificationGorm.GetOtpVerificationDetailsByUserPID(c, pid)

	if err != nil {
		return err
	}

	return err
}
