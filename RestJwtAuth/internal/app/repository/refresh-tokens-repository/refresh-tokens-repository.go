package refresh_tokens_repository

import (
	refresh_token_dto "RestJwtAuth/internal/app/models/refresh-token-dto"
	"RestJwtAuth/internal/pkg/app"
	"go.mongodb.org/mongo-driver/mongo"
)

// TODO: implement RefreshTokensCRUDRepository methods
type RefreshTokensRepository struct {
	collection mongo.Collection
}

func New(client *mongo.Client) *RefreshTokensRepository {
	c := *client.Database(app.DbName).Collection(app.RefreshTokensCollection)
	return &RefreshTokensRepository{collection: c}
}

// TODO: implement Save method
func (r *RefreshTokensRepository) Save(dto *refresh_token_dto.RefreshTokenDTO) error {
	return nil
}

// TODO: implement Find method
func (r *RefreshTokensRepository) Find(subject string) (*refresh_token_dto.RefreshTokenDTO, error) {
	return nil, nil
}
