package request

import "split_it_backend/businesses/users"

type Register struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (reg *Register) ToDomain() *users.DomainUser{
	return &users.DomainUser{
		Email: reg.Email,
		Password: reg.Password,
	}
}