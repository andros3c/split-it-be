package users

import (
	"context"
	"split_it_backend/businesses/users"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(gormDb *gorm.DB)users.Repository{
	return &UserRepository{
		db: gormDb,
	}
}

func(repo *UserRepository)Register(domain users.DomainUser,ctx context.Context)(users.DomainUser,error){
	userDb := FromDomain(domain)

	err:= repo.db.Create(&userDb).Error

	if err != nil{
		return users.DomainUser{},err
	}
	return userDb.ToDomain(),err
}

func (repo *UserRepository)FindUserByEmailOrUsername(domain users.DomainUser,ctx context.Context)(users.DomainUser,error){
	userDb := FromDomain(domain)

	err := repo.db.Where("username = ? OR email = ?",userDb.Username,userDb.Email).First(&userDb).Error
	if err == gorm.ErrRecordNotFound{
		return userDb.ToDomain(),nil
	}else if err == nil{
		return users.DomainUser{},gorm.ErrDuplicatedKey
	}else{
		return users.DomainUser{},err
	}
}

func (repo *UserRepository)Login(domain users.DomainUser,ctx context.Context)(users.DomainUser,error){
	userDb := FromDomain(domain)
	res := repo.db.Where("username = ? OR email = ?",userDb.Username,userDb.Email).First(&userDb)

	if res.Error != nil{
		return users.DomainUser{},res.Error
	}
	return userDb.ToDomain(),nil
}