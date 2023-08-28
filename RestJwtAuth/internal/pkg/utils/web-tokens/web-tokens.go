package web_tokens

import (
	"RestJwtAuth/internal/app/models/user"
	user_details "RestJwtAuth/internal/app/models/user-details"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
	"time"
)

const (
	RefreshTokenHeader   = "Refresh-Token"
	AccessTokenHeader    = "Access-Token"
	TokenPrefix          = "Bearer "
	AccessTokenLifetime  = time.Minute * 5
	RefreshTokenLifetime = time.Hour * 48
)

type JWTParsable interface {
	Subject() string
}

func isAuthenticationToken(token string) bool {
	return strings.HasPrefix(token, TokenPrefix)
}

func HasRefreshToken(header http.Header) bool {
	refresh := header.Get(RefreshTokenHeader)
	return isAuthenticationToken(refresh)
}

func HasAccessToken(header http.Header) bool {
	access := header.Get(AccessTokenHeader)
	return isAuthenticationToken(access)
}

func BuildAccessJWT(u JWTParsable, secret string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub":             u.Subject(),
		"generation_time": time.Now().Add(AccessTokenLifetime).Format(time.Layout),
	})
	return t.SignedString([]byte(secret))
}

func BuildRefreshToken(u JWTParsable, secret string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub":             u.Subject(),
		"generation_time": time.Now().Add(RefreshTokenLifetime).Format(time.Layout),
	})
	return t.SignedString([]byte(secret))
}

func BuildAccessRefreshPair(u *user.User, secret string) (access, refresh string, err error) {
	access, err = BuildAccessJWT(user_details.Of(u), secret)
	if err != nil {
		return "", "", err
	}
	refresh, err = BuildRefreshToken(user_details.Of(u), secret)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}
