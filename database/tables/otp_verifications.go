package tables

import (
	"time"
)

type OtpVerifications struct {
	ID                     int    `gorm:"column:otp_id;primaryKey;autoIncrement"`
	PID                    string `gorm:"column:otp_pid;unique;not null;type:varchar(40)"`
	UserPID                string `gorm:"column:user_pid;not null;type:varchar(40)"`
	OtpType                string `gorm:"column:otp_type;type:varchar(20)"`
	OtpValue               string `gorm:"column:otp_value;type:varchar(20)"`
	OtpStatus              string `gorm:"column:otp_status;type:varchar(20)"`
	VerificationRetryCount int    `gorm:"column:verification_retry_count;type:int;default:0"`
	ResendCount            int    `gorm:"column:resend_count;type:int;default:0"`
	IsDeleted              bool   `gorm:"column:is_deleted;not null;default:false"`
	IsSandbox              bool   `gorm:"column:is_sandbox;not null;default:false"`
	Metadata               JSONB  `gorm:"column:metadata;type:json"`
	CreatedAt              time.Time
	UpdatedAt              time.Time
}
