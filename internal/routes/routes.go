package routes

import (
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/config"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, cfg *config.Config, db *gorm.DB) {
	userHandler := handler.NewUserHandler(r, cfg, db)

    api := r.Group("/api/v1")
    {
        user := api.Group("/user")
        {
            user.POST("/register", userHandler.Register)
            user.POST("/login", userHandler.Login)
            user.POST("/refresh", userHandler.Refresh)
        }
    }
}