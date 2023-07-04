package app

import (
	"github.com/BearTS/go-gin-monolith/app/middleware"
	"github.com/BearTS/go-gin-monolith/controllers/v1/user"
	"github.com/BearTS/go-gin-monolith/database"
	"github.com/BearTS/go-gin-monolith/dbops/gorm/otp_verifications"
	"github.com/BearTS/go-gin-monolith/dbops/gorm/users"
	"github.com/BearTS/go-gin-monolith/services/authsvc"
	"github.com/BearTS/go-gin-monolith/services/usersvc"
	"github.com/gin-gonic/gin"
)

func MapURL() {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	gormDB, _ := database.Connection()
	usersGorm := users.Gorm(gormDB)
	otpVerificationsGorm := otp_verifications.Gorm(gormDB)

	authsvc := authsvc.Handler()
	userSvc := usersvc.Handler(usersGorm, otpVerificationsGorm, authsvc)

	// Handlers
	userHandler := user.Handler(userSvc)

	v1 := router.Group("/v1")

	users := v1.Group("/users")
	{
		users.POST("/send-otp", userHandler.SendOTP)
		users.POST("/verify-otp", userHandler.VerifyOTP)
		users.POST("/resend-otp", userHandler.ResendOTP)
	}

	err := router.Run()
	if err != nil {
		panic(err.Error() + "MapURL router not able to run")
	}
}
