package user

import (
	"authapp/dto/request"
)

type Repository interface {
	Persist(u *User) (*User, error)

	GetUserByUserPass(phonenumber, password string) (*User, error)
	GetUserByPhonenumber(phonenumber string) (*User, error)
}

type Service interface {
	CreateUserIfNotAny(req request.CreateUserRequest) (*User, error)

	Login(phonenumber, password string) (*User, string, error)
}
