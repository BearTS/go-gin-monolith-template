package authsvc

// /* -------------------------------------------------------------------------- */
// /*                                Token Request                               */
// /* -------------------------------------------------------------------------- */
// type TokenReq struct {
// 	Type        string      `json:"type"`
// 	CustomerID  string      `json:"customer_id"`
// 	TransfersID string      `json:"transfers_id"`
// 	Metadata    interface{} `json:"metadata"`
// }

/* -------------------------------------------------------------------------- */
/*                               Token Response                               */
/* -------------------------------------------------------------------------- */
type TokenRes struct {
	Type             string       `json:"type"`
	UserID           string       `json:"user_id,omitempty"`
	AdminID         string       `json:"admin_id,omitempty"`
	IsPhoneAvailable bool         `json:"is_phone_available,omitempty"`
	Metadata         *interface{} `json:"metadata"`
	AccesssTokenPID  string       `json:"access_token_pid"`
	AccesssToken     string       `json:"access_token"`
	AccessTokenExp   int64        `json:"access_token_expiry"`
	RefreshTokenPID  string       `json:"refresh_token_pid"`
	RefreshToken     string       `json:"refresh_token"`
	RefreshTokenExp  int64        `json:"refresh_token_expiry"`
}

/* -------------------------------------------------------------------------- */
/*                                  Auth Data                                 */
/* -------------------------------------------------------------------------- */
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	TokenUuid    string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

/* -------------------------------------------------------------------------- */
/*                           Access Secret For Auth                           */
/* -------------------------------------------------------------------------- */
type AccessSecretOrg struct {
	AccessKey string
	SecretKey string
}

/* -------------------------------------------------------------------------- */
/*                            Onboarding Token Req                            */
/* -------------------------------------------------------------------------- */
type TokenReq struct {
	Type     string      `json:"type"`
	UserID   string      `json:"user_id"`
	Metadata interface{} `json:"metadata"`
}
