package auth

import (
	"OrderApp/persistency/table"
)

func (a *ServiceImpl) CreateUser(dto CreateUserCommand) error {
	hash, e := hashPassword(dto.Password)
	if e != nil {
		return e
	}

	user := table.User{
		Name:     dto.Name,
		Role:     dto.Role,
		Email:    dto.Email,
		Password: hash,
	}
	return a.userPersistency.CreateUser(user)
}

type CreateUserCommand struct {
	Name     string
	Role     string
	Email    string
	Password string
}
