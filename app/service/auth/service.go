package auth

import (
	"OrderApp/common/postgresql/table"
	"OrderApp/persistency"
)

type Service interface {
	GetUser(id string) (*table.User, error)
	CreateUser(dto CreateUserCommand) error
	Login(email, password string) (string, error)
}

type ServiceImpl struct {
	userPersistency persistency.UserPersistency
}

func NewService(userPersistency persistency.UserPersistency) Service {
	return &ServiceImpl{
		userPersistency: userPersistency,
	}
}
