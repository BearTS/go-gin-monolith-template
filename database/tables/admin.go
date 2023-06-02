package tables

import "time"

type Admins struct {
	ID        int    `gorm:"column:admin_id;primaryKey;autoIncrement"`
	PID       string `gorm:"column:admin_pid;unique;not null;type:varchar(40)"`
	Name      string `gorm:"column:name;not null;type:varchar(100)"`
	Email     string `gorm:"column:email;not null;type:varchar(40)"`
	Password  []byte `gorm:"column:password;type:bytea"`
	IsSandbox bool   `gorm:"column:is_sandbox;not null;default:false"`
	IsDeleted bool   `gorm:"column:is_deleted;not null;default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
