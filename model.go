package account

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/weihongguo/gglmm"
)

// Authenticationable --
type Authenticationable interface {
	CheckPassword(password string) error
	GenerateJWT(jwtExpires int64, jwtSecret string) (string, jwt.StandardClaims, error)
	GenerateAuthenticationInfo() (*AuthenticationInfo, error)
}

// Account --
type Account struct {
	gglmm.Model
	Password  string `json:"-"`
	Nickname  string `json:"nickname"`
	AvatarURL string `json:"avatarURL"`
}

var (
	// AccountStatuses --
	AccountStatuses = []gglmm.Status{gglmm.StatusValid, gglmm.StatusFrozen, gglmm.StatusInvalid}
)

// CheckPassword --
func (account Account) CheckPassword(password string) error {
	return gglmm.BcryptCompareHashAndPassword(account.Password, password)
}

func (account Account) generateJWT(userType string, expires int64, jwtSecret string) (string, jwt.StandardClaims, error) {
	jwtUser := &JWTUser{}
	jwtUser.UserType = userType
	jwtUser.UserID = account.ID
	jwtUser.UserName = account.Nickname

	return GenerateJWTToken(jwtUser, expires, jwtSecret)
}

// GenerateAuthenticationInfo --
func (account Account) GenerateAuthenticationInfo() (*AuthenticationInfo, error) {
	info := &AuthenticationInfo{}
	info.ID = account.ID
	info.Nickname = account.Nickname
	info.AvatarURL = account.AvatarURL
	return info, nil
}
