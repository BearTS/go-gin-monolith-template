package admin

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
	CreateAdmin(ctx *gin.Context, admin tables.Admins) (tables.Admins, error)
	GetAdminByPID(ctx *gin.Context, pid string) (tables.Admins, error)
	GetAdminByEmail(ctx *gin.Context, email string) (tables.Admins, error)
	GetAdminDetails(ctx *gin.Context) (tables.Admins, error)
	UpdateAdminDetails(ctx *gin.Context, admin tables.Admins) (tables.Admins, error)
}

/* -------------------------------------------------------------------------- */
/*                                   Handler                                  */
/* -------------------------------------------------------------------------- */
func Gorm(gormDB *gorm.DB) *adminGormImpl {
	return &adminGormImpl{
		DB: gormDB,
	}
}

type adminGormImpl struct {
	DB *gorm.DB
}

/* -------------------------------------------------------------------------- */
/*                                   Methods                                  */
/* -------------------------------------------------------------------------- */

func (r *adminGormImpl) CreateAdmin(ctx *gin.Context, admin tables.Admins) (tables.Admins, error) {
	admin.PID = utils.UUIDWithPrefix(constants.Prefix.ADMIN)

	err := r.DB.Session(&gorm.Session{}).Create(&admin).Error
	if err != nil {
		return admin, errors.Wrap(err, "[adminGormImpl][CreateAdmin]")
	}
	return admin, nil
}

func (r *adminGormImpl) GetAdminByPID(ctx *gin.Context, pid string) (tables.Admins, error) {
	var admin tables.Admins

	err := r.DB.Session(&gorm.Session{}).Where("admin_pid = ?", pid).
		Scopes(dbops.DeletedScopes(ctx)).
		First(&admin).Error

	if err != nil {
		return admin, errors.Wrap(err, "[adminGormImpl][GetAdminByPID]")
	}
	return admin, nil
}

// Get Runner By Email
func (r *adminGormImpl) GetAdminByEmail(ctx *gin.Context, email string) (tables.Admins, error) {
	var admin tables.Admins

	err := r.DB.Session(&gorm.Session{}).Where("email = ?", email).
		Scopes(dbops.DeletedScopes(ctx)).
		First(&admin).Error

	if err != nil {
		return admin, errors.Wrap(err, "[adminGormImpl][GetAdminByEmail]")
	}
	return admin, nil
}

// Get Runner details
func (r *adminGormImpl) GetAdminDetails(ctx *gin.Context) (tables.Admins, error) {
	var admin tables.Admins

	authData, err := utils.GetAuthData(ctx)
	if err != nil {
		return admin, errors.Wrap(err, "[authData][GetAdminDetails]")
	}

	err = r.DB.Session(&gorm.Session{}).Where("admin_pid = ?", authData.AdminPID).
		Scopes(dbops.DeletedScopes(ctx)).
		Scopes(dbops.SandboxScopes(ctx)).
		First(&admin).Error

	if err != nil {
		return admin, errors.Wrap(err, "[adminGormImpl][GetAdminDetails]")
	}

	return admin, nil
}

// Update Runner details
func (r *adminGormImpl) UpdateAdminDetails(ctx *gin.Context, admin tables.Admins) (tables.Admins, error) {
	authData, err := utils.GetAuthData(ctx)
	if err != nil {
		return admin, errors.Wrap(err, "[authData][UpdateAdminDetails]")
	}

	err = r.DB.Session(&gorm.Session{}).Where("admin_pid = ?", authData.AdminPID).
		Scopes(dbops.DeletedScopes(ctx)).
		Scopes(dbops.SandboxScopes(ctx)).
		Updates(&admin).Error

	if err != nil {
		return admin, errors.Wrap(err, "[adminGormImpl][UpdateAdminDetails]")
	}

	return admin, nil
}

