package dbops

import (
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// database common scopes will come here
func SandboxCustomerDeleted(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	authData, _ := utils.GetAuthData(c)
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("is_sandbox = ?", authData.Sandbox).
			Where("user_pid = ?", authData.UserPID).
			Where("is_deleted = ?", false)
	}
}

// onboarding scope
func SandboxDeleted(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	authData, _ := utils.GetAuthData(c)
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("is_sandbox = ?", authData.Sandbox).
			Where("is_deleted = ?", false)
	}
}

func UserScopes(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	authData, _ := utils.GetAuthData(c)
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("user_pid = ?", authData.UserPID)
	}
}

func SandboxScopes(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	authData, _ := utils.GetAuthData(c)
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("is_sandbox = ?", authData.Sandbox)
	}
}

func DeletedScopes(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("is_deleted = ?", false)
	}
}

func ActiveScopes(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("is_active = ?", true)
	}
}

func LatestScopes(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("is_latest = ?", true)
	}
}

func RunnerScopes(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	authData, _ := utils.GetAuthData(c)
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("runner_pid = ?", authData.UserPID)
	}
}
