package users

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
	CreateUser(ctx *gin.Context, user tables.Users) (tables.Users, error)
	GetUserDetailsByPID(ctx *gin.Context, PID string) (tables.Users, error)
	UpdateUser(ctx *gin.Context, user tables.Users) (tables.Users, error)
	GetUserDetails(ctx *gin.Context) (tables.Users, error)
	GetUserDetailsByEmail(ctx *gin.Context, email string) (tables.Users, error)
}

/* -------------------------------------------------------------------------- */
/*                                   Handler                                  */
/* -------------------------------------------------------------------------- */
func Gorm(gormDB *gorm.DB) *usersGormImpl {
	return &usersGormImpl{
		DB: gormDB,
	}
}

type usersGormImpl struct {
	DB *gorm.DB
}

/* -------------------------------------------------------------------------- */
/*                                   Methods                                  */
/* -------------------------------------------------------------------------- */

// get or read example
func (r *usersGormImpl) GetUserDetailsByPID(ctx *gin.Context, PID string) (tables.Users, error) {
	var user tables.Users

	db := r.DB.Session(&gorm.Session{})
	result := db.Where("user_pid = ?", PID).
		Scopes(dbops.DeletedScopes(ctx)).
		Take(&user)

	err := result.Error
	if err != nil {
		return user, err
	}
	return user, err
}

// create example
func (r *usersGormImpl) CreateUser(ctx *gin.Context, user tables.Users) (tables.Users, error) {
	user.PID = utils.UUIDWithPrefix(constants.Prefix.USER)

	err := r.DB.Session(&gorm.Session{}).Create(&user).Error
	if err != nil {
		return user, errors.Wrap(err, "[usersGormImpl][CreateUser]")
	}
	return user, nil
}

// update example
func (r *usersGormImpl) UpdateUser(ctx *gin.Context, user tables.Users) (tables.Users, error) {
	authData, err := utils.GetAuthData(ctx)
	if err != nil {
		return user, errors.Wrap(err, "[GetUserDetail][GetAuthData]")
	}
	user.PID = authData.UserPID
	db := r.DB.Session(&gorm.Session{})
	result := db.Where("user_pid = ?", user.PID).
		Scopes(dbops.SandboxScopes(ctx)).
		Scopes(dbops.DeletedScopes(ctx)).
		Updates(user)
	err = result.Error
	if err != nil {
		return user, err
	}
	return user, err
}

// GetUserDetails method
func (r *usersGormImpl) GetUserDetails(ctx *gin.Context) (tables.Users, error) {
	var user tables.Users
	authData, err := utils.GetAuthData(ctx)
	if err != nil {
		return user, errors.Wrap(err, "[GetUserDetails][GetAuthData]")
	}

	db := r.DB.Session(&gorm.Session{})
	result := db.Where("user_pid = ?", authData.UserPID).
		Scopes(dbops.SandboxScopes(ctx)).
		Scopes(dbops.DeletedScopes(ctx)).
		Take(&user)

	err = result.Error
	if err != nil {
		return user, errors.Wrap(err, "[GetUserDetails]")
	}
	return user, err
}

func (r *usersGormImpl) GetUserDetailsByEmail(ctx *gin.Context, email string) (tables.Users, error) {
	var user tables.Users

	db := r.DB.Session(&gorm.Session{})
	result := db.Where("user_email = ?", email).
		Scopes(dbops.DeletedScopes(ctx)).
		Take(&user)

	err := result.Error
	if err != nil {
		return user, err
	}
	return user, err
}
