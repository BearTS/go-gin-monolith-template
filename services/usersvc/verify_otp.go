package usersvc

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/BearTS/go-gin-monolith/constants"
	"github.com/BearTS/go-gin-monolith/database/tables"
	"github.com/BearTS/go-gin-monolith/services/authsvc"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (g *UserSvcImpl) VerifyOTP(c *gin.Context, req VerifyOTPReq) (utils.BaseResponse, tables.Users, error) {
	var baseRes utils.BaseResponse
	var res tables.Users
	var err error

	// Add initial Response
	baseRes.Success = false
	baseRes.StatusCode = http.StatusInternalServerError
	baseRes.Message = "Internal Server Error"

	// search user by pid in otp verification
	otpVerification, err := g.otpVerificationGorm.GetOtpVerificationDetailsByUserPID(c, req.UserPID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			baseRes.StatusCode = http.StatusNotFound
			baseRes.Message = constants.Errors.VerifyOTP.InvalidOTP
			return baseRes, res, err
		}
		return baseRes, res, err
	}

	if otpVerification.OtpValue != req.Otp {
		baseRes.StatusCode = http.StatusUnauthorized
		baseRes.Message = constants.Errors.VerifyOTP.InvalidOTP
		return baseRes, res, err
	}

	/* -------------------------------------------------------------------------- */
	/*                              Expiry Time Check                             */
	/* -------------------------------------------------------------------------- */

	// Check expiry
	currentTime := time.Now()
	// expires at is time when otp was created + 5 minutes
	expiresAt := otpVerification.CreatedAt.Add(time.Minute * 5)
	if currentTime.After(expiresAt) {

		// Update the existing otp
		otpVerification.OtpStatus = constants.OtpStatuses.EXPIRED
		go func() {
			_, err = g.otpVerificationGorm.UpdateOtpVerification(c, otpVerification)
			if err != nil {
				log.Println("[Critical][Should Not Happen] Error while updating otp status to expired")
			}
		}()

		baseRes.StatusCode = http.StatusUnauthorized
		baseRes.Message = constants.Errors.VerifyOTP.InvalidOTP
		return baseRes, res, err
	}

	/* -------------------------------------------------------------------------- */
	/*                              Update Otp Status                             */
	/* -------------------------------------------------------------------------- */

	otpVerification.OtpStatus = constants.OtpStatuses.VERIFIED
	otpVerification, err = g.otpVerificationGorm.UpdateOtpVerification(c, otpVerification)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			baseRes.StatusCode = http.StatusNotFound
			baseRes.Message = "Internal Server Error"

			return baseRes, res, err
		}
		return baseRes, res, err
	}

	// search user by pid in users
	user, err := g.usersGorm.GetUserDetailsByPID(c, req.UserPID)
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
	// 	// Have token generation with only Type = first time
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
