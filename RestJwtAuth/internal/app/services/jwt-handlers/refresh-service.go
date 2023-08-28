package jwt_handlers

import refresh_token_dto "RestJwtAuth/internal/app/models/refresh-token-dto"

type RefreshTokensRepository interface {
	Save(dto *refresh_token_dto.RefreshTokenDTO) error
	Find(subject string) (*refresh_token_dto.RefreshTokenDTO, error)
}

type RefreshService struct {
	rtRepository RefreshTokensRepository
}

func NewRefreshService(repo RefreshTokensRepository) *RefreshService {
	return &RefreshService{rtRepository: repo}
}

// TODO: implement Refresh func
func (r *RefreshService) Refresh(access, refresh string) (newAccess, newRefresh string, err error) {
	//check access and refresh tokens generation type compliance

	//get refresh token from db

	//check hash of refresh param with token from db

	//create new access and refresh token

	return "", "", nil
}
