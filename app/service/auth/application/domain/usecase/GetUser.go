package usecase

import (
	"OrderApp/service/auth/application/domain/model"
	"OrderApp/service/auth/application/port/out"
)

type GetUser struct {
	userPersistencePort out.UserPersistencePort
}

func NewGetUser(userPersistencePort out.UserPersistencePort) *GetUser {
	return &GetUser{userPersistencePort: userPersistencePort}
}

func (c GetUser) GetUser(id string) (*model.User, error) {
	return c.userPersistencePort.GetUser(id)
}
