package users

import (
	"context"

	"split_it_backend/app/middleware"
	"split_it_backend/businesses"
	"split_it_backend/drivers/database/helpers/encrypt"
	"time"

	"gorm.io/gorm"
)

type UserUseCase struct {
	repo Repository
	ctx  time.Duration
	jwt	 *middleware.ConfigJWT
}

func NewUseCase(userRepo Repository,contextTimeout time.Duration,configJWT middleware.ConfigJWT)UseCase{

	return &UserUseCase{
		repo: userRepo,
		ctx: contextTimeout,
		jwt: &configJWT,
	}
}

func (usecase *UserUseCase)Register(domain DomainUser,ctx context.Context)(DomainUser,error){
	if domain.Email == ""{
		return DomainUser{},businesses.ErrEmailEmpty
	}
	if domain.Username == ""{
		return DomainUser{},businesses.ErrUsernameEmpty
	}
	if domain.Password == ""{
		return DomainUser{},businesses.ErrPassEmpty
	}


	_,err := usecase.repo.FindUserByEmailOrUsername(domain,ctx)

	if err == gorm.ErrDuplicatedKey{
		return DomainUser{},businesses.ErrDuplicateUser
	} else if err != nil{
		return DomainUser{},err
	}

	
		domain.Password, err = encrypt.Hash(domain.Password)
		if err != nil{
			return DomainUser{},businesses.ErrInternalServer
		}
		user,errRegister := usecase.repo.Register(domain,ctx)
	
		if errRegister != nil{
			return DomainUser{},err
		}
	
		return user,nil
	
	
}

func (usecase *UserUseCase)Login(domain DomainUser,ctx context.Context)(DomainUser,error){
	// if domain.Username == ""{
	// 	return DomainUser{},businesses.ErrUsernameEmpty
	// }
	if domain.Password == ""{
		return DomainUser{},businesses.ErrPassEmpty
	}
	
	userData,err := usecase.repo.Login(domain,ctx)
	if err == gorm.ErrRecordNotFound {
		return DomainUser{},businesses.ErrUserNotFound
	}
	if err != nil{
		return DomainUser{},err
	}

	if !encrypt.ValidateHash(domain.Password,userData.Password){
		return DomainUser{},businesses.ErrWrongPassword
	}
	userData.Token = usecase.jwt.GenererateToken(uint(userData.Id))
	 

return userData,nil
}