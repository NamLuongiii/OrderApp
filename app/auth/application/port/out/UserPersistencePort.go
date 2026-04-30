package out

import "OrderApp/auth/application/domain/model"

type UserPersistencePort interface {
	GetUser(id string) (*model.User, error)
	CreateUser(user *model.User) error
	GetUserByEmail(email string) (*model.User, error)
}
