package routes

import (
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/config"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/handler"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, cfg *config.Config, db *gorm.DB) {
	userHandler := handler.NewUserHandler(cfg, db)
	postHandler := handler.NewPostHandler(cfg, db)

	r.Static("/uploads", "./uploads")

	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", userHandler.Register)
			auth.POST("/login", userHandler.Login)
			auth.POST("/refresh", userHandler.Refresh)
		}
		posts := api.Group("/posts")
		{
			posts.GET("", postHandler.GetAll)
			posts.GET("/:id", postHandler.GetById)

			protected := posts.Group("")
			protected.Use(middleware.AuthMiddleware(cfg.JWT_SECRET))
			{
				protected.POST("/create", postHandler.Create)
				protected.PUT("/update/:id", postHandler.Update)
				protected.DELETE("/delete/:id", postHandler.Delete)
			}
		}
	}
}