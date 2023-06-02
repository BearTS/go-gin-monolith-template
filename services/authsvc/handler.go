package authsvc

import (
	"github.com/BearTS/go-gin-monolith/models"
	"github.com/BearTS/go-gin-monolith/utils"
	"github.com/gin-gonic/gin"
)

type authSvcImpl struct{}

// interface.
type Interface interface {
	GenerateToken(c *gin.Context, req TokenReq) (utils.BaseResponse, TokenRes, error)
	CreateToken(tokenAuthData models.AuthData) (*TokenDetails, error)
	ValidateToken(signedToken string) error
}

/* -------------------------------------------------------------------------- */
/*                              AUTHSVC HANDLER                               */
/* -------------------------------------------------------------------------- */
func Handler() *authSvcImpl {
	return &authSvcImpl{}
}
