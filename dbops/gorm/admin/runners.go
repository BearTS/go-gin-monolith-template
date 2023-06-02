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
	CreateRunner(ctx *gin.Context, runner tables.Admins) (tables.Admins, error)
	GetRunnerByPID(ctx *gin.Context, pid string) (tables.Admins, error)
	GetRunnerByEmail(ctx *gin.Context, email string) (tables.Admins, error)
	GetRunnerDetails(ctx *gin.Context) (tables.Admins, error)
	UpdateRunnerDetails(ctx *gin.Context, runner tables.Admins) (tables.Admins, error)
	GetAllActiveRunners(ctx *gin.Context) ([]tables.Admins, error)
}

/* -------------------------------------------------------------------------- */
/*                                   Handler                                  */
/* -------------------------------------------------------------------------- */
func Gorm(gormDB *gorm.DB) *runnersGormImpl {
	return &runnersGormImpl{
		DB: gormDB,
	}
}

type runnersGormImpl struct {
	DB *gorm.DB
}

/* -------------------------------------------------------------------------- */
/*                                   Methods                                  */
/* -------------------------------------------------------------------------- */

func (r *runnersGormImpl) CreateRunner(ctx *gin.Context, runner tables.Admins) (tables.Admins, error) {
	runner.PID = utils.UUIDWithPrefix(constants.Prefix.ADMIN)

	err := r.DB.Session(&gorm.Session{}).Create(&runner).Error
	if err != nil {
		return runner, errors.Wrap(err, "[runnersGormImpl][CreateRunner]")
	}
	return runner, nil
}

func (r *runnersGormImpl) GetRunnerByPID(ctx *gin.Context, pid string) (tables.Admins, error) {
	var runner tables.Admins

	err := r.DB.Session(&gorm.Session{}).Where("admin_pid = ?", pid).
		Scopes(dbops.DeletedScopes(ctx)).
		First(&runner).Error

	if err != nil {
		return runner, errors.Wrap(err, "[runnersGormImpl][GetRunnerByPID]")
	}
	return runner, nil
}

// Get Runner By Email
func (r *runnersGormImpl) GetRunnerByEmail(ctx *gin.Context, email string) (tables.Admins, error) {
	var runner tables.Admins

	err := r.DB.Session(&gorm.Session{}).Where("email = ?", email).
		Scopes(dbops.DeletedScopes(ctx)).
		First(&runner).Error

	if err != nil {
		return runner, errors.Wrap(err, "[runnersGormImpl][GetRunnerByEmail]")
	}
	return runner, nil
}

// Get Runner details
func (r *runnersGormImpl) GetRunnerDetails(ctx *gin.Context) (tables.Admins, error) {
	var runner tables.Admins

	authData, err := utils.GetAuthData(ctx)
	if err != nil {
		return runner, errors.Wrap(err, "[authData][GetRunnerDetails]")
	}

	err = r.DB.Session(&gorm.Session{}).Where("admin_pid = ?", authData.AdminPID).
		Scopes(dbops.DeletedScopes(ctx)).
		Scopes(dbops.SandboxScopes(ctx)).
		First(&runner).Error

	if err != nil {
		return runner, errors.Wrap(err, "[runnersGormImpl][GetRunnerDetails]")
	}

	return runner, nil
}

// Update Runner details
func (r *runnersGormImpl) UpdateRunnerDetails(ctx *gin.Context, runner tables.Admins) (tables.Admins, error) {
	authData, err := utils.GetAuthData(ctx)
	if err != nil {
		return runner, errors.Wrap(err, "[authData][UpdateRunnerDetails]")
	}

	err = r.DB.Session(&gorm.Session{}).Where("admin_pid = ?", authData.AdminPID).
		Scopes(dbops.DeletedScopes(ctx)).
		Scopes(dbops.SandboxScopes(ctx)).
		Updates(&runner).Error

	if err != nil {
		return runner, errors.Wrap(err, "[runnersGormImpl][UpdateRunnerDetails]")
	}

	return runner, nil
}

func (r *runnersGormImpl) GetAllActiveRunners(ctx *gin.Context) ([]tables.Admins, error) {
	var runners []tables.Admins

	_, err := utils.GetAuthData(ctx)
	if err != nil {
		return runners, errors.Wrap(err, "[authData][UpdateRunnerDetails]")
	}

	err = r.DB.Session(&gorm.Session{}).Where("is_active = ?", true).
		Scopes(dbops.DeletedScopes(ctx)).
		Scopes(dbops.SandboxScopes(ctx)).
		Find(&runners).Error

	if err != nil {
		return runners, errors.Wrap(err, "[runnersGormImpl][UpdateRunnerDetails]")
	}

	return runners, nil
}
