package response

import (
	"split_it_backend/businesses/users"
	"time"

	"gorm.io/gorm"
)

type UserResponse struct {
	Id                   uint32           `json:"id"`
	Username             string         `json:"username"`
	Email                string         `json:"email"`
	Token				 string			`json:"token,omitempty"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
	DeletedAt            gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain users.DomainUser)UserResponse{
	return UserResponse{
		Id: domain.Id,
		Username: domain.Username,
		Email: domain.Email,
		Token: domain.Token,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}