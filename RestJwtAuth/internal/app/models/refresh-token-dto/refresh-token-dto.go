package refresh_token_dto

type RefreshTokenDTO struct {
	Subject   string `bson:"subject,omitempty"`
	HashToken string `bson:"hash_token,omitempty"`
	Blocked   bool   `bson:"blocked,"`
}

func New(ownerUsername, hashToken string) *RefreshTokenDTO {
	return &RefreshTokenDTO{
		Subject:   ownerUsername,
		HashToken: hashToken,
		Blocked:   false,
	}
}
