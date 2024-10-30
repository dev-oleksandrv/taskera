package request

import "github.com/dev-oleksandrv/taskera/internal/model/domain"

type UserRegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (r *UserRegisterRequest) ToDomainUser() domain.User {
	return domain.User{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
	}
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
