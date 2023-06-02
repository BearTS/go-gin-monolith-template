package tables

import "time"

type Admins struct {
	ID              int    `gorm:"column:admin_id;primaryKey;autoIncrement"`
	PID             string `gorm:"column:admin_pid;unique;not null;type:varchar(40)"`
	Name            string `gorm:"column:name;not null;type:varchar(100)"`
	Phone           string `gorm:"column:phone;not null;type:varchar(10)"`
	Email           string `gorm:"column:email;not null;type:varchar(40)"`
	Password        []byte `gorm:"column:password;type:bytea"`
	AvailablePoints int    `gorm:"column:available_points;not null;default:0"`
	RedeemedPoints  int    `gorm:"column:redeemed_points;not null;default:0"`
	Photo           string `gorm:"column:photo;type:varchar(100)"`
	IsActive        bool   `gorm:"column:is_active;not null;default:true"`
	IsSandbox       bool   `gorm:"column:is_sandbox;not null;default:false"`
	IsDeleted       bool   `gorm:"column:is_deleted;not null;default:false"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
