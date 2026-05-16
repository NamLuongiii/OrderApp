package auth

import (
	"OrderApp/common/msg"
	"errors"
)

func (a *ServiceImpl) Login(email, password string) (string, error) {
	user, e := a.userPersistency.GetUserByEmail(email)
	if e != nil {
		return "", errors.New(msg.InvalidCredentials)
	}
	if !checkPasswordCorrect(password, user.Password) {
		return "", errors.New(msg.InvalidCredentials)
	}
	token, e := generateJwtToken(user.ID, Role(user.Role))
	if e != nil {
		return "", errors.New(msg.InternalServerError)
	}
	return token, nil
}
