package seeds

import (
	"github.com/BearTS/go-gin-monolith/database/tables"
	"gorm.io/gorm"
)

func Admin(db *gorm.DB) error {
	// seed 1
	err := db.Create(&tables.Admins{
		PID:       "adm_89b2f151110644cb84b0912bb8ce7885",
		Name:      "Test Admin",
		Email:     "admin@Admin",
		Phone:     "1234567890",
		IsDeleted: false,
		IsSandbox: false,
	}).Error

	if err != nil {
		return err
	}

	return err
}
