package businesses

import "errors"

var (
	ErrEmailEmpty = errors.New("Email Empty")
	ErrUsernameEmpty = errors.New("Username Empty")
	ErrPassEmpty = errors.New("Password Empty")
	ErrDuplicateUser = errors.New("username or email already used")	
	ErrInternalServer = errors.New("something's wrong, contact administrator")
	ErrWrongPassword = errors.New("wrong password")
	ErrUserNotFound = errors.New("user not found")
)