package main

import (
	"github.com/dev-oleksandrv/taskera/internal/config"
	"github.com/dev-oleksandrv/taskera/internal/database"
	"github.com/dev-oleksandrv/taskera/internal/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Init(cfg)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	g := r.Group("/api")

	router.UserRouter(g, db)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	if err := r.Run(":" + cfg.Server.Port); err != nil {
		panic(err)
	}
}
