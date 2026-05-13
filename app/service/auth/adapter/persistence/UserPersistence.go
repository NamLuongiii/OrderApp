package persistence

import (
	"OrderApp/common/postgresql/table"
	model2 "OrderApp/service/auth/application/domain/model"

	"gorm.io/gorm"
)

type UserPersistence struct {
	db *gorm.DB
}

func NewUserPersistence(db *gorm.DB) *UserPersistence {
	return &UserPersistence{db: db}
}

func (p *UserPersistence) CreateUser(user *model2.User) error {
	persistenceUser := table.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     string(user.Role),
	}

	e := p.db.Create(&persistenceUser).Error

	return e
}

func (p *UserPersistence) GetUser(id string) (*model2.User, error) {
	persistenceUser := table.User{}
	e := p.db.Where("id = ?", id).First(&persistenceUser).Error
	if e != nil {
		return nil, e
	}
	user := model2.User{
		ID:        persistenceUser.ID,
		Name:      persistenceUser.Name,
		Email:     persistenceUser.Email,
		Password:  persistenceUser.Password,
		Role:      model2.Role(persistenceUser.Role),
		CreatedAt: persistenceUser.CreatedAt.Unix(),
		UpdatedAt: persistenceUser.UpdatedAt.Unix(),
	}
	return &user, nil
}

func (p *UserPersistence) GetUserByEmail(email string) (*model2.User, error) {
	persistenceUser := table.User{}
	e := p.db.Where("email = ?", email).First(&persistenceUser).Error
	if e != nil {
		return nil, e
	}
	user := model2.User{
		ID:        persistenceUser.ID,
		Name:      persistenceUser.Name,
		Email:     persistenceUser.Email,
		Password:  persistenceUser.Password,
		Role:      model2.Role(persistenceUser.Role),
		CreatedAt: persistenceUser.CreatedAt.Unix(),
		UpdatedAt: persistenceUser.UpdatedAt.Unix(),
	}
	return &user, nil
}
