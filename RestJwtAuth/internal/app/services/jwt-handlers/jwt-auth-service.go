package jwt_handlers

import (
	refresh_token_dto "RestJwtAuth/internal/app/models/refresh-token-dto"
	"RestJwtAuth/internal/app/models/user"
	"RestJwtAuth/internal/pkg/utils/hashcoding"
	web_tokens "RestJwtAuth/internal/pkg/utils/web-tokens"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UsersRepository interface {
	Create(u user.User) error
	Read(username string) (*user.User, error)
	Update(username string, u user.User) (*user.User, error)
	Delete(username string) error
}

type JWTAuthService struct {
	usersRepository UsersRepository
	rtRepository    RefreshTokensRepository
	secretKey       string
}

func NewJWTAuthService(usersRepository UsersRepository, rtRepository RefreshTokensRepository, secretKey string) *JWTAuthService {
	s := &JWTAuthService{
		usersRepository: usersRepository,
		rtRepository:    rtRepository,
		secretKey:       secretKey,
	}
	return s
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *JWTAuthService) Authorize(authData map[string]string) (access, refresh string, err error) {
	u, err := s.usersRepository.Read(authData["username"])
	if err != nil {
		return "", "", err
	}
	bytes, err := hashcoding.BcryptRefreshToken(authData["password"])
	if err != nil {
		return "", "", fmt.Errorf("password encryption error")
	}
	c := string(bytes)
	if u.HashPassword() != c {
		return "", "", fmt.Errorf("password is not correct")

	}
	access, refresh, err = web_tokens.BuildAccessRefreshPair(u, s.secretKey)

	rtDTO := refresh_token_dto.New(u.Username(), refresh)
	s.rtRepository.Save(rtDTO)

	return access, refresh, err
}
