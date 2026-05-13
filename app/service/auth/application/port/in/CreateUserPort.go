package in

import (
	"OrderApp/service/auth/application/domain/model"
)

type CreateUserPort interface {
	CreateUser(*model.User) error
}
