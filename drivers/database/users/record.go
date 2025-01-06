package users

import (
	"split_it_backend/businesses/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint32    `gorm:"primaryKey"`
	Username  string    
	Email     string    
	Password  string   
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FromDomain(domain users.DomainUser)User{
	return User{
		Id: domain.Id,
		Username: domain.Username,
		Email: domain.Email,
		Password: domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}

func (user User)ToDomain()users.DomainUser{
	return users.DomainUser{
		Id: user.Id,
		Username: user.Username,
		Email: user.Email,
		Password: user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}