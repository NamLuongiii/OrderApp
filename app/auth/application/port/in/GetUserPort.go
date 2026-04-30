package in

import "OrderApp/auth/application/domain/model"

type GetUserPort interface {
	GetUser(userId string) (*model.User, error)
}
