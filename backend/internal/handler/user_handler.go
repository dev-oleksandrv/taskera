package handler

import (
	"errors"
	"github.com/dev-oleksandrv/taskera/internal/model/http/request"
	"github.com/dev-oleksandrv/taskera/internal/model/http/response"
	"github.com/dev-oleksandrv/taskera/internal/service"
	"github.com/dev-oleksandrv/taskera/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) Register(ctx *gin.Context) {
	var input request.UserRegisterRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		handleError(ctx, http.StatusBadRequest, utils.MapErrorToValidationErrors(err))
		return
	}
	user := input.ToDomainUser()
	if err := h.userService.Register(&user); err != nil {
		if errors.Is(err, service.ErrEmailAlreadyExist) {
			handleError(ctx, http.StatusBadRequest, map[string]string{
				"email": "Email already exists",
			})
			return
		}

		handleError(ctx, http.StatusInternalServerError, map[string]string{
			"server": "Something went wrong",
		})
		return
	}

	handleSuccess(ctx, http.StatusCreated, response.UserRegisterResponse{
		User: user.ToDto(),
	})
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var input request.UserLoginRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		handleError(ctx, http.StatusBadRequest, utils.MapErrorToValidationErrors(err))
		return
	}

	user, err := h.userService.Login(input.Email, input.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidUserPassword) {
			handleError(ctx, http.StatusBadRequest, map[string]string{
				"password": "Password is invalid",
			})
			return
		}

		if errors.Is(err, service.ErrEmailNotExist) {
			handleError(ctx, http.StatusBadRequest, map[string]string{
				"email": "Email does not exist",
			})
			return
		}

		handleError(ctx, http.StatusInternalServerError, map[string]string{
			"server": "Something went wrong",
		})
		return
	}

	handleSuccess(ctx, http.StatusOK, response.UserLoginResponse{
		User:  user.ToDto(),
		Token: utils.GenerateJWTToken(user.ID, user.Email),
	})
}

func handleError(ctx *gin.Context, statusCode int, errors map[string]string) {
	ctx.AbortWithStatusJSON(statusCode, response.ErrorResponse{
		Errors: errors,
	})
}

func handleSuccess(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode, response.SuccessResponse{
		Data: data,
	})
}
