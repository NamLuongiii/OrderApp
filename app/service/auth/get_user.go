package auth

import (
	"OrderApp/persistency/table"
)

func (a *ServiceImpl) GetUser(id string) (*table.User, error) {
	return a.userPersistency.GetUser(id)
}
