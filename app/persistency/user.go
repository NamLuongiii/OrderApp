package persistency

import (
	"OrderApp/common/postgresql/table"

	"gorm.io/gorm"
)

type UserPersistency interface {
	CreateUser(user table.User) error
	GetUser(id string) (*table.User, error)
	GetUserByEmail(email string) (*table.User, error)
}

type UserPersistenceImpl struct {
	db *gorm.DB
}

func NewUserPersistence(db *gorm.DB) UserPersistency {
	return &UserPersistenceImpl{db: db}
}

func (p *UserPersistenceImpl) CreateUser(user table.User) error {
	e := p.db.Create(&user).Error
	return e
}

func (p *UserPersistenceImpl) GetUser(id string) (*table.User, error) {
	user := table.User{}
	e := p.db.Where("id = ?", id).First(&user).Error
	if e != nil {
		return nil, e
	}

	return &user, nil
}

func (p *UserPersistenceImpl) GetUserByEmail(email string) (*table.User, error) {
	user := table.User{}
	e := p.db.Where("email = ?", email).First(&user).Error
	if e != nil {
		return nil, e
	}
	return &user, nil
}
