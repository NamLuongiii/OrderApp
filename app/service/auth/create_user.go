package auth

import "OrderApp/common/postgresql/table"

func (a *ServiceImpl) CreateUser(dto CreateUserCommand) error {
	user := table.User{
		Name:     dto.Name,
		Role:     dto.Role,
		Email:    dto.Email,
		Password: dto.Password,
	}
	return a.userPersistency.CreateUser(user)
}

type CreateUserCommand struct {
	Name     string
	Role     string
	Email    string
	Password string
}
