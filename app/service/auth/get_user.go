package auth

import "OrderApp/common/postgresql/table"

func (a *ServiceImpl) GetUser(id string) (*table.User, error) {
	return a.userPersistency.GetUser(id)
}
