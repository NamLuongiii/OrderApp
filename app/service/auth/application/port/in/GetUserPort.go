package in

import (
	"OrderApp/service/auth/application/domain/model"
)

type GetUserPort interface {
	GetUser(userId string) (*model.User, error)
}
