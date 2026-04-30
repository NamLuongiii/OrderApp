package usecase

import (
	"OrderApp/auth/application/domain/model"
	"OrderApp/auth/application/port/out"
)

type CreateUserUseCase struct {
	userPersistencePort out.UserPersistencePort
}

func NewCreateUserUseCase(userPersistencePort out.UserPersistencePort) *CreateUserUseCase {
	return &CreateUserUseCase{userPersistencePort: userPersistencePort}
}

func (u *CreateUserUseCase) CreateUser(user *model.User) error {
	e := u.userPersistencePort.CreateUser(user)
	return e
}
