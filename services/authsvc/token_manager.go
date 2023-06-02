package authsvc

import (
	"fmt"
	"time"

	"github.com/BearTS/go-gin-monolith/config"
	"github.com/BearTS/go-gin-monolith/constants"
	"github.com/BearTS/go-gin-monolith/models"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

func (t *authSvcImpl) CreateToken(tokenAuthData models.AuthData) (*TokenDetails, error) {
	var err error
	// access token expiry
	atexp := time.Now().Add(time.Minute * 30) //expires after 30 min

	// refresh token expiry
	rtexp := time.Now().Add(time.Hour * 24 * 7) // expires after 7 days

	// set authdata
	tokenAuthData.SessionPID = utils.UUIDWithPrefix(constants.Prefix.SESSION)
	tokenAuthData.RegisteredClaims.Issuer = "sample-issuers"
	tokenAuthData.RegisteredClaims.IssuedAt = &jwt.NumericDate{time.Now()}

	td := &TokenDetails{}
	td.AtExpires = atexp.Unix()
	td.TokenUuid = utils.UUIDWithPrefix("tk")

	td.RtExpires = rtexp.Unix()
	td.RefreshUuid = td.TokenUuid + "++" + tokenAuthData.UserPID

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["session_pid"] = tokenAuthData.SessionPID
	atClaims["type"] = tokenAuthData.Type
	atClaims["sandbox"] = tokenAuthData.Sandbox
	atClaims["exp"] = td.AtExpires
	atClaims["iss"] = tokenAuthData.RegisteredClaims.Issuer
	atClaims["iat"] = tokenAuthData.RegisteredClaims.IssuedAt

	if tokenAuthData.AdminPID != "" {
		atClaims["admin_pid"] = tokenAuthData.AdminPID
	} else if tokenAuthData.UserPID != "" {
		atClaims["user_pid"] = tokenAuthData.UserPID
	} else {
		return nil, errors.New("invalid token auth data")
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(config.Token.AccessSecret))
	if err != nil {
		return nil, errors.Wrap(err, "[CreateToken][AccessToken]")
	}

	//Creating Refresh Token
	td.RtExpires = rtexp.Unix()
	td.RefreshUuid = td.TokenUuid + "++" + tokenAuthData.UserPID

	//set auth data
	tokenAuthData.SessionPID = utils.UUIDWithPrefix(constants.Prefix.SESSION)
	tokenAuthData.RegisteredClaims.Issuer = "tez"
	tokenAuthData.RegisteredClaims.IssuedAt = &jwt.NumericDate{time.Now()}

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_session_pid"] = td.RefreshUuid
	rtClaims["type"] = "refresh"
	rtClaims["sandbox"] = tokenAuthData.Sandbox
	rtClaims["exp"] = td.AtExpires
	rtClaims["issuer"] = tokenAuthData.RegisteredClaims.Issuer
	rtClaims["issued_at"] = tokenAuthData.RegisteredClaims.IssuedAt

	if tokenAuthData.AdminPID != "" {
		rtClaims["admin_pid"] = tokenAuthData.AdminPID
	}

	if tokenAuthData.UserPID != "" {
		rtClaims["user_pid"] = tokenAuthData.UserPID
	}

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	td.RefreshToken, err = rt.SignedString([]byte(config.Token.RefreshSecret))
	if err != nil {
		return nil, errors.Wrap(err, "[CreateToken][RefreshToken]")
	}
	return td, nil
}

func (s *authSvcImpl) ValidateToken(signedToken string) error {
	accessSecretKey := []byte(config.Token.AccessSecret)
	token, err := jwt.ParseWithClaims(signedToken, &models.AuthData{}, func(t *jwt.Token) (interface{}, error) {
		return accessSecretKey, nil
	})
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*models.AuthData)
	if !ok {
		err = errors.New("couldn't parse claims")
		return err
	}
	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return err
	}

	// extra verification
	_, err = jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Token.AccessSecret), nil
	})
	if err != nil {
		return err
	}

	return err
}
