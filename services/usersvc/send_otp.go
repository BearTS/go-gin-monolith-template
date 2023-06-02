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

func (g *UserSvcImpl) SendOTP(c *gin.Context, req SendOTPReq) (utils.BaseResponse, tables.Users, error) {
	var baseRes utils.BaseResponse
	var res tables.Users
	var err error

	// Add initial Response
	baseRes.Success = false
	baseRes.StatusCode = http.StatusInternalServerError
	baseRes.Message = "Internal Server Error"

	// search user by email
	res, err = g.usersGorm.GetUserDetailsByEmail(c, req.Email)
	if err == nil {
		// Check for existing otp
		_, err := g.otpVerificationGorm.GetOtpVerificationDetailsByUserPID(c, res.PID)
		if err == nil {
			baseRes.StatusCode = http.StatusConflict
			baseRes.Message = "OTP already sent"
			return baseRes, res, err
		}

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return baseRes, res, err
		}
		// Generate a new OTP
		otp := utils.GenerateOtp(6)

		// Save
		var otpVerification tables.OtpVerifications
		otpVerification.OtpValue = otp
		otpVerification.UserPID = res.PID
		otpVerification.OtpType = "email"
		otpVerification.OtpStatus = constants.OtpStatuses.PENDING

		_, err = g.otpVerificationGorm.CreateNewOTPVerification(c, otpVerification)

		body := "Your OTP is " + otp
		subject := "TEZ: your otp is " + otp
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

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return baseRes, res, err
	}

	// Create a new user
	var user tables.Users
	user.Email = req.Email

	// Create a new user
	res, err = g.usersGorm.CreateUser(c, user)
	if err != nil {
		return baseRes, res, err
	}

	// Generate a new OTP
	otp := utils.GenerateOtp(6)

	var otpVerification tables.OtpVerifications
	otpVerification.UserPID = res.PID
	otpVerification.OtpValue = otp
	otpVerification.OtpType = "email"
	otpVerification.OtpStatus = constants.OtpStatuses.PENDING

	_, err = g.otpVerificationGorm.CreateNewOTPVerification(c, otpVerification)
	if err != nil {
		return baseRes, res, err
	}

	body := "Your OTP is " + otp
	subject := "TEZ: your otp is " + otp
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
	baseRes.Message = "User Created Successfully"

	return baseRes, res, err
}
