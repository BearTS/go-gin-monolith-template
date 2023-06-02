package seeds

import (
	"github.com/BearTS/go-gin-monolith/database/tables"
	"gorm.io/gorm"
)

func Devices(db *gorm.DB) error {
	// Seed 1
	err := db.Create(&tables.Devices{
		PID:         "dev_89b2f151110644cb84b0912bb8ce7886",
		UserPID:     "usr_gh67fhgCU123jkhfdsjkdfhkds",
		DeviceToken: "fcm_token_1",
	}).Error

	if err != nil {
		return err
	}

	return err
}
