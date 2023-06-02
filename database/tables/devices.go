package tables

import "time"

type Devices struct {
	ID          int    `gorm:"column:devices_id;primaryKey;autoIncrement"`
	PID         string `gorm:"column:devices_pid;unique;not null;type:varchar(40)"`
	UserPID     string `gorm:"column:user_pid;not null;type:varchar(40)"`
	DeviceToken string `gorm:"column:device_token;not null;type:varchar(255)"`
	IsDeleted   bool   `gorm:"column:is_deleted;not null;default:false"`
	IsSandbox   bool   `gorm:"column:is_sandbox;not null;default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
