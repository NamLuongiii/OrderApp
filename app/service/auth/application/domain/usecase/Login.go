package usecase

import (
	model2 "OrderApp/service/auth/application/domain/model"
	"OrderApp/service/auth/application/port/out"
	"errors"
	"fmt"
)

type LoginUseCase struct {
	userPersistencePort out.UserPersistencePort
}

func NewLoginUseCase(userPersistencePort out.UserPersistencePort) *LoginUseCase {
	return &LoginUseCase{userPersistencePort: userPersistencePort}
}

func (l LoginUseCase) Login(email, password string) (string, error) {
	user, e := l.userPersistencePort.GetUserByEmail(email)
	if e != nil {
		return "", e
	}

	if model2.ValidatePassword(password, user.Password) == false {
		return "", errors.New("user or password is wrong")
	}

	token, e := model2.GenerateJwtToken(user.ID, user.Role)

	fmt.Println(token, e)
	if e != nil {
		return "", e
	}

	return token, nil
}
