package user_details

import "RestJwtAuth/internal/app/models/user"

type UserDetails struct {
	subject string
}

func Of(u *user.User) *UserDetails {
	return &UserDetails{subject: u.Username()}
}

func (d *UserDetails) Subject() string {
	return d.subject
}
