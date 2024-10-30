package response

import (
	"github.com/dev-oleksandrv/taskera/internal/model/domain"
)

type UserRegisterResponse struct {
	User domain.UserDto `json:"user"`
}

type UserLoginResponse struct {
	User  domain.UserDto `json:"user"`
	Token string         `json:"token"`
}
