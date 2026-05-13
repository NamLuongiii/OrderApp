package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID        string
	Role      Role
	Email     string
	Password  string
	Name      string
	CreatedAt int64
	UpdatedAt int64
}

func NewUser(
	id string,
	role Role,
	email string,
	password string,
	name string,
	createdAt int64,
	updatedAt int64,
) *User {
	return &User{
		ID:        id,
		Role:      role,
		Email:     email,
		Password:  password,
		Name:      name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func NewUserWithoutId(
	role Role,
	email string,
	password string,
	name string,
) (*User, error) {
	p, e := HashPassword(password)
	if e != nil {
		return nil, e
	}
	return &User{
		Role:     role,
		Email:    email,
		Password: p,
		Name:     name,
	}, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func ValidatePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
