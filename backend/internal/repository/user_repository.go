package repository

import (
	"errors"
	"github.com/dev-oleksandrv/taskera/internal/model/domain"
	"gorm.io/gorm"
)

var (
	ErrInvalidUserEmail    = errors.New("invalid user email")
	ErrInvalidUserPassword = errors.New("invalid user password")
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(user *domain.User) error {
	tx := r.db.Begin()
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
