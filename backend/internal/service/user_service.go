package service

import (
	"errors"
	"github.com/dev-oleksandrv/taskera/internal/model/domain"
	"github.com/dev-oleksandrv/taskera/internal/repository"
	"github.com/dev-oleksandrv/taskera/internal/utils"
	"strings"
)

var (
	ErrEmailNotExist       = errors.New("email not exist")
	ErrEmailAlreadyExist   = errors.New("email already exist")
	ErrInvalidUserPassword = errors.New("invalid user password")
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (s *UserService) Register(user *domain.User) error {
	if err := s.userRepository.Create(user); err != nil {
		if strings.Contains(err.Error(), "uni_users_email") {
			return ErrEmailAlreadyExist
		}
		return err
	}
	return nil
}

func (s *UserService) Login(email, password string) (*domain.User, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return nil, ErrEmailNotExist
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, ErrInvalidUserPassword
	}
	return user, nil
}
