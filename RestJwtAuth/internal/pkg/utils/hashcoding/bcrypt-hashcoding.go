package hashcoding

import "golang.org/x/crypto/bcrypt"

func BcryptUserPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func BcryptRefreshToken(refreshToken string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}
