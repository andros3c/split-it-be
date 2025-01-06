package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type DomainUser struct {
	Id        uint32    `gorm:"primaryKey" json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Token	  string	`json:"token"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UseCase interface{
	Register(domain DomainUser,ctx context.Context)(DomainUser,error)
	Login(domain DomainUser,ctx context.Context)(DomainUser,error)
}

type Repository interface{
	Register(domain DomainUser,ctx context.Context)(DomainUser,error)
	FindUserByEmailOrUsername(domain DomainUser,ctx context.Context)(DomainUser,error)
	Login(domain DomainUser,ctx context.Context)(DomainUser,error)
}