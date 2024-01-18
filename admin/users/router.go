package users

import (
	"github.com/JingruiLea/ad_boost/dal/user_dal"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"time"
)

type UserData struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Roles        []string  `json:"roles"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	Expires      time.Time `json:"expires"`
}

var MyAuthenticator = &jwt.GinJWTMiddleware{
	Realm:            "test zone",
	SigningAlgorithm: "",
	Key:              []byte("secret key"),
	KeyFunc:          nil,
	Timeout:          time.Hour,
	MaxRefresh:       time.Hour,
	Authenticator:    Authenticator,
	Authorizator:     nil,
	PayloadFunc:      nil,
	Unauthorized:     nil,
	LoginResponse: func(c *gin.Context, code int, token string, expires time.Time) {
		userData, ok := c.Get("user")
		if !ok {
			return
		}
		user := userData.(*UserData)
		user.AccessToken = token
		user.Expires = expires
		c.JSON(200, map[string]interface{}{
			"code": 0,
			"msg":  "success",
			"data": map[string]interface{}{
				"access_token":  token,
				"refresh_token": token,
				"expires":       expires,

				"id":       user.ID,
				"username": user.Username,
				"roles":    user.Roles,
			},
		})
	},
	LogoutResponse:        nil,
	RefreshResponse:       nil,
	IdentityHandler:       nil,
	IdentityKey:           "",
	TokenLookup:           "",
	TokenHeadName:         "",
	TimeFunc:              nil,
	HTTPStatusMessageFunc: nil,
	PrivKeyFile:           "",
	PrivKeyBytes:          nil,
	PubKeyFile:            "",
	PrivateKeyPassphrase:  "",
	PubKeyBytes:           nil,
	SendCookie:            false,
	CookieMaxAge:          0,
	SecureCookie:          false,
	CookieHTTPOnly:        false,
	CookieDomain:          "",
	SendAuthorization:     false,
	DisabledAbort:         false,
	CookieName:            "",
	CookieSameSite:        0,
	ParseOptions:          nil,
}

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Authenticator(c *gin.Context) (interface{}, error) {
	var loginVals login
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	username := loginVals.Username
	password := loginVals.Password

	user, err := user_dal.GetUserByLogin(c, username, password)
	if err != nil {
		return nil, err
	}
	if user != nil {
		userData := &UserData{
			ID:       int(user.ID), // 示例 ID
			Username: user.Username,
			Roles:    []string{"admin"},
		}
		c.Set("user", userData)
		return nil, nil
	}
	return nil, jwt.ErrFailedAuthentication
}
