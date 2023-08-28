package user

type User struct {
	username     string `json:"username" bson:"username"`
	hashPassword string `json:"-" bson:"hash_password"`
}

func New(username, hashPassword string) *User {
	return &User{
		username:     username,
		hashPassword: hashPassword,
	}
}

func (u *User) Username() string {
	return u.username
}

func (u *User) HashPassword() string {
	return u.hashPassword
}
