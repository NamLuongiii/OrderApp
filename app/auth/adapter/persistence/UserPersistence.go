package persistence

import (
	"OrderApp/auth/application/domain/model"
	"OrderApp/common/postgresql/table"

	"gorm.io/gorm"
)

type UserPersistence struct {
	db *gorm.DB
}

func NewUserPersistence(db *gorm.DB) *UserPersistence {
	return &UserPersistence{db: db}
}

func (p *UserPersistence) CreateUser(user *model.User) error {
	persistenceUser := table.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     string(user.Role),
	}

	e := p.db.Create(&persistenceUser).Error

	return e
}

func (p *UserPersistence) GetUser(id string) (*model.User, error) {
	persistenceUser := table.User{}
	e := p.db.Where("id = ?", id).First(&persistenceUser).Error
	if e != nil {
		return nil, e
	}
	user := model.User{
		ID:        persistenceUser.ID,
		Name:      persistenceUser.Name,
		Email:     persistenceUser.Email,
		Password:  persistenceUser.Password,
		Role:      model.Role(persistenceUser.Role),
		CreatedAt: persistenceUser.CreatedAt.Unix(),
		UpdatedAt: persistenceUser.UpdatedAt.Unix(),
	}
	return &user, nil
}

func (p *UserPersistence) GetUserByEmail(email string) (*model.User, error) {
	persistenceUser := table.User{}
	e := p.db.Where("email = ?", email).First(&persistenceUser).Error
	if e != nil {
		return nil, e
	}
	user := model.User{
		ID:        persistenceUser.ID,
		Name:      persistenceUser.Name,
		Email:     persistenceUser.Email,
		Password:  persistenceUser.Password,
		Role:      model.Role(persistenceUser.Role),
		CreatedAt: persistenceUser.CreatedAt.Unix(),
		UpdatedAt: persistenceUser.UpdatedAt.Unix(),
	}
	return &user, nil
}
