package router

import (
	"github.com/dev-oleksandrv/taskera/internal/handler"
	"github.com/dev-oleksandrv/taskera/internal/repository"
	"github.com/dev-oleksandrv/taskera/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouter(router *gin.RouterGroup, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(*userRepository)
	userHandler := handler.NewUserHandler(*userService)

	group := router.Group("/user")

	group.POST("/register", userHandler.Register)
	group.POST("/login", userHandler.Login)
}
