package request

import "split_it_backend/businesses/users"

type Login struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (reg *Login) ToDomain() *users.DomainUser{
	return &users.DomainUser{
		Username: reg.Username,
		Email: reg.Email,
		Password: reg.Password,
	}
}