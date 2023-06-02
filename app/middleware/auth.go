package app

import (
	"github.com/BearTS/go-gin-monolith/constants"
	"github.com/BearTS/go-gin-monolith/services/authsvc"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
)

func TokenAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(401, gin.H{"error": "request does not contain an access token"})
			ctx.Abort()
			return
		}
		// Remove the Bearer prefix and take the token
		tokenHandler := authsvc.Handler()
		err := tokenHandler.ValidateToken(token)
		if err != nil {
			ctx.JSON(401, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func CheckIfCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authData, err := utils.GetAuthData(ctx)
		if err != nil {
			ctx.JSON(401, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		switch authData.Type {
		case constants.TokenTypes.USER:
			ctx.Next()
		default:
			ctx.JSON(401, gin.H{"error": "user is not a customer"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func CheckIfAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authData, err := utils.GetAuthData(ctx)
		if err != nil {
			ctx.JSON(401, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		if authData.Type != constants.TokenTypes.ADMIN {
			ctx.JSON(401, gin.H{"error": "user is not a admin"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
