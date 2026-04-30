package in

import "OrderApp/auth/application/domain/model"

type CreateUserPort interface {
	CreateUser(*model.User) error
}
