package otp_verifications

import (
	"github.com/BearTS/go-gin-monolith/constants"
	"github.com/BearTS/go-gin-monolith/database/tables"
	"github.com/BearTS/go-gin-monolith/dbops"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

/* -------------------------------------------------------------------------- */
/*                                  Interface                                 */
/* -------------------------------------------------------------------------- */
type GormInterface interface {
	CreateNewOTPVerification(ctx *gin.Context, otpVerification tables.OtpVerifications) (tables.OtpVerifications, error)
	CreateOTPVerification(ctx *gin.Context, otpVerification tables.OtpVerifications) (tables.OtpVerifications, error)
	GetOtpVerificationDetailsByPID(ctx *gin.Context, PID string) (tables.OtpVerifications, error)
	GetOtpVerificationDetailsByUserPID(ctx *gin.Context, UserPID string) (tables.OtpVerifications, error)
	UpdateOtpVerification(ctx *gin.Context, otpVerification tables.OtpVerifications) (tables.OtpVerifications, error)
}

/* -------------------------------------------------------------------------- */
/*                                   Handler                                  */
/* -------------------------------------------------------------------------- */
func Gorm(gormDB *gorm.DB) *otpVerificationsGormImpl {
	return &otpVerificationsGormImpl{
		DB: gormDB,
	}
}

type otpVerificationsGormImpl struct {
	DB *gorm.DB
}

/* -------------------------------------------------------------------------- */
/*                                   Methods                                  */
/* -------------------------------------------------------------------------- */

// CreateNewOTPVerification method for unauthenticated users
func (r *otpVerificationsGormImpl) CreateNewOTPVerification(ctx *gin.Context, otpVerification tables.OtpVerifications) (tables.OtpVerifications, error) {

	otpVerification.PID = utils.UUIDWithPrefix(constants.Prefix.OTPVERIFICATION)

	err := r.DB.Session(&gorm.Session{}).Create(&otpVerification).Error
	if err != nil {
		return otpVerification, errors.Wrap(err, "[otpVerificationsGormImpl][CreateNewOTPVerification]")
	}
	return otpVerification, nil
}

// CreateOTPVerification method
func (r *otpVerificationsGormImpl) CreateOTPVerification(ctx *gin.Context, otpVerification tables.OtpVerifications) (tables.OtpVerifications, error) {

	authData, _ := utils.GetAuthData(ctx)

	otpVerification.UserPID = authData.UserPID
	otpVerification.IsSandbox = authData.Sandbox

	err := r.DB.Session(&gorm.Session{}).Create(&otpVerification).Error
	if err != nil {
		return otpVerification, errors.Wrap(err, "[otpVerificationsGormImpl][CreateOTPVerification]")
	}
	return otpVerification, nil
}

// GetOtpVerificationDetailsByUserPID method
func (r *otpVerificationsGormImpl) GetOtpVerificationDetailsByUserPID(ctx *gin.Context, UserPID string) (tables.OtpVerifications, error) {
	var otpVerification tables.OtpVerifications

	db := r.DB.Session(&gorm.Session{})

	result := db.Where("user_pid = ?", UserPID).
		Where("otp_status = ?", constants.OtpStatuses.PENDING).
		Scopes(dbops.DeletedScopes(ctx)).
		Order("created_at DESC").
		Take(&otpVerification)

	err := result.Error
	if err != nil {
		return otpVerification, err
	}
	return otpVerification, err
}

// GetOtpVerificationDetailsByPID method
func (r *otpVerificationsGormImpl) GetOtpVerificationDetailsByPID(ctx *gin.Context, PID string) (tables.OtpVerifications, error) {
	var otpVerification tables.OtpVerifications

	db := r.DB.Session(&gorm.Session{})
	result := db.Where("otp_pid = ?", PID).
		Where("otp_status = ?", constants.OtpStatuses.PENDING).
		Scopes(dbops.SandboxScopes(ctx)).
		Scopes(dbops.DeletedScopes(ctx)).
		Scopes(dbops.UserScopes(ctx)).
		Take(&otpVerification)

	err := result.Error
	if err != nil {
		return otpVerification, err
	}
	return otpVerification, err
}

// UpdateOtpVerification method
func (r *otpVerificationsGormImpl) UpdateOtpVerification(ctx *gin.Context, otpVerification tables.OtpVerifications) (tables.OtpVerifications, error) {
	db := r.DB.Session(&gorm.Session{})
	result := db.Where("otp_pid = ?", otpVerification.PID).
		Updates(otpVerification)
	err := result.Error
	if err != nil {
		return otpVerification, err
	}
	return otpVerification, err
}
