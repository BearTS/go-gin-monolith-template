package usersvc

import (
	"errors"
	"log"
	"net/http"

	"github.com/BearTS/go-gin-monolith/constants"
	"github.com/BearTS/go-gin-monolith/database/tables"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (g *UserSvcImpl) ResendOTP(c *gin.Context, req ResendOTPReq) (utils.BaseResponse, tables.Users, error) {
	var baseRes utils.BaseResponse
	var res tables.Users
	var err error

	// Add initial Response
	baseRes.Success = false
	baseRes.StatusCode = http.StatusInternalServerError
	baseRes.Message = "Internal Server Error"

	// search user by email
	res, err = g.usersGorm.GetUserDetailsByEmail(c, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			baseRes.StatusCode = http.StatusNotFound
			baseRes.Message = "User not found"
			return baseRes, res, err
		}
		return baseRes, res, err
	}

	// Check for existing otp
	otp, err := g.otpVerificationGorm.GetOtpVerificationDetailsByUserPID(c, res.PID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return baseRes, res, err
	}

	// Update the existing otp
	otp.OtpStatus = constants.OtpStatuses.EXPIRED
	otp, err = g.otpVerificationGorm.UpdateOtpVerification(c, otp)
	if err != nil {
		return baseRes, res, err
	}

	// generate new otp
	otpValue := utils.GenerateOtp(6)

	var otpVerification tables.OtpVerifications
	otpVerification.OtpValue = otpValue
	otpVerification.UserPID = res.PID
	otpVerification.OtpType = "email"
	otpVerification.OtpStatus = constants.OtpStatuses.PENDING

	_, err = g.otpVerificationGorm.CreateNewOTPVerification(c, otpVerification)

	body := "Your OTP is " + otpValue
	subject := "TEZ: your otp is " + otpValue
	toAddress := req.Email

	go func() {
		err := utils.SendEmail(subject, body, []string{toAddress}, nil, nil, nil)
		if err != nil {
			log.Println("Failed to send email: ", err)
		}
	}()
	// Add Success Response
	baseRes.Success = true
	baseRes.StatusCode = http.StatusOK
	baseRes.Message = "OTP sent Successfully"

	return baseRes, res, err
}
