package user

import (
	"net/http"

	"github.com/BearTS/go-gin-monolith/merrors"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*                                Authorization                               */
/* -------------------------------------------------------------------------- */

/* -------------------------------- Send Otp -------------------------------- */
func (h *userHandler) SendOTP(c *gin.Context) {

	req, err := validateSendOTPReq(c)
	if err != nil {
		merrors.Validation(c, err.Error())
		return
	}

	baseRes, res, err := h.usersvc.SendOTP(c, req)
	if err != nil {
		merrors.InternalServer(c, baseRes.Message)
		return
	}

	if baseRes.StatusCode != http.StatusOK {
		merrors.InternalServer(c, baseRes.Message)
		return
	}

	finalRes := sendOtpTransformer(res)

	utils.ReturnJSONStruct(c, finalRes)
}

/* ------------------------------- Resend OTP ------------------------------- */
func (h *userHandler) ResendOTP(c *gin.Context) {
	req, err := validateResendOTPReq(c)
	if err != nil {
		merrors.Validation(c, err.Error())
		return
	}

	baseRes, res, err := h.usersvc.ResendOTP(c, req)
	if err != nil {
		merrors.InternalServer(c, baseRes.Message)
		return
	}

	if baseRes.StatusCode != http.StatusOK {
		merrors.InternalServer(c, baseRes.Message)
		return
	}

	finalRes := sendOtpTransformer(res)

	utils.ReturnJSONStruct(c, finalRes)
}

/* ------------------------------- Verify Otp ------------------------------- */
func (h *userHandler) VerifyOTP(c *gin.Context) {

	req, err := validateVerifyOTPReq(c)
	if err != nil {
		merrors.Validation(c, err.Error())
		return
	}

	baseRes, _, err := h.usersvc.VerifyOTP(c, req)
	if err != nil {
		merrors.InternalServer(c, baseRes.Message)
		return
	}

	if baseRes.StatusCode != http.StatusOK {
		merrors.InternalServer(c, baseRes.Message)
		return
	}

	finalRes := baseRes

	utils.ReturnJSONStruct(c, finalRes)
}
