package devices

import (
	"github.com/BearTS/go-gin-monolith/database/tables"
	"github.com/BearTS/go-gin-monolith/dbops"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type GormInterface interface {
	AddNewDevice(c *gin.Context, device tables.Devices) (tables.Devices, error)
	GetDeviceByToken(c *gin.Context, Token string) (tables.Devices, error)
	GetDevicesByUserPID(c *gin.Context, userPID string) ([]tables.Devices, error)
}

/* -------------------------------------------------------------------------- */
/*                                   Handler                                  */
/* -------------------------------------------------------------------------- */
func Gorm(gormDB *gorm.DB) *deviceGormImpl {
	return &deviceGormImpl{
		DB: gormDB,
	}
}

type deviceGormImpl struct {
	DB *gorm.DB
}

func (r *deviceGormImpl) AddNewDevice(c *gin.Context, device tables.Devices) (tables.Devices, error) {
	device.PID = utils.UUIDWithPrefix("dev")
	err := r.DB.Session(&gorm.Session{}).Create(&device).Error
	if err != nil {
		return device, errors.Wrap(err, "[deviceGormImpl][AddNewDevices]")
	}
	return device, nil
}

func (r *deviceGormImpl) GetDeviceByToken(c *gin.Context, Token string) (tables.Devices, error) {
	var device tables.Devices
	err := r.DB.Session(&gorm.Session{}).
		Where("device_token = ?", Token).
		Scopes(dbops.DeletedScopes(c)).
		Scopes(dbops.SandboxScopes(c)).
		First(&device).Error

	if err != nil {
		return device, errors.Wrap(err, "[deviceGormImpl][GetDeviceByToken]")
	}
	return device, nil
}

func (r *deviceGormImpl) GetDevicesByUserPID(c *gin.Context, userPID string) ([]tables.Devices, error) {
	var devices []tables.Devices
	err := r.DB.Session(&gorm.Session{}).
		Where("user_pid = ?", userPID).
		Scopes(dbops.DeletedScopes(c)).
		Scopes(dbops.SandboxScopes(c)).
		Find(&devices).Error

	if err != nil {
		return devices, errors.Wrap(err, "[deviceGormImpl][GetDevicesByUserPID]")
	}
	return devices, nil
}
