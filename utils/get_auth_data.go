package utils

import (
	"github.com/BearTS/go-gin-monolith/config"
	"github.com/BearTS/go-gin-monolith/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

func GetAuthData(ctx *gin.Context) (*models.AuthData, error) {
	signedToken := ctx.GetHeader("Authorization")
	return GetAuthDataFromToken(signedToken)
}

func GetRefreshTokenData(ctx *gin.Context) (*models.AuthData, error) {
	signedToken := ctx.GetHeader("Authorization")
	return GetRefreshTokenDataFromToken(signedToken)
}

func GetAuthDataFromToken(signedToken string) (*models.AuthData, error) {
	var res *models.AuthData
	jwtKey := []byte(config.Token.AccessSecret)
	token, err := jwt.ParseWithClaims(signedToken, &models.AuthData{}, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return res, err
	}
	res, ok := token.Claims.(*models.AuthData)
	if !ok {
		err = errors.New("couldn't parse claims")
		return res, err
	}
	return res, err
}

func GetRefreshTokenDataFromToken(signedToken string) (*models.AuthData, error) {
	var res *models.AuthData
	jwtKey := []byte(config.Token.RefreshSecret)
	token, err := jwt.ParseWithClaims(signedToken, &models.AuthData{}, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return res, err
	}
	res, ok := token.Claims.(*models.AuthData)
	if !ok {
		err = errors.New("couldn't parse claims")
		return res, err
	}
	return res, err
}
