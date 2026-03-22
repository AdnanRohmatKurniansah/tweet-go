package handler

import (
	"net/http"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/config"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/dto"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/repository"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/service"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	api *gin.Engine
	userService service.UserService
}

func NewUserHandler(api *gin.Engine, cfg *config.Config, db *gorm.DB) *UserHandler {
	repo := repository.NewUserRepository(db)
	userSvc := service.NewService(cfg, repo)

	return &UserHandler{
		api: api,
		userService: userSvc,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
    var req dto.RegisterRequest
    if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := utils.FormatValidationError(err)

		c.JSON(422, gin.H{
			"message": "Validation failed",
			"errors":  validationErrors,
		})
		return
	}

    res, err := h.userService.Register(req); 
    
    if err != nil {
        utils.ErrorMessage(c, 500, "Failed to register user", err.Error())
        return
    }

    utils.SuccessMessage(c, http.StatusCreated, "Register success", res)
}

func (h *UserHandler) Login(c *gin.Context) {
    var req dto.LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        utils.ErrorMessage(c, 422, "Failed to process request", err.Error())
        return
    }

    res, err := h.userService.Login(req)
    if err != nil {
        utils.ErrorMessage(c, 401, "Failed to login", err.Error())
        return
    }

	utils.SuccessMessage(c, 200, "Login success", res)
}

func (h *UserHandler) Refresh(c *gin.Context) {
    var req dto.RefreshRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        utils.ErrorMessage(c, 422, "Failed to process request", err.Error())
        return
    }

    res, err := h.userService.Refresh(req)
    if err != nil {
        utils.ErrorMessage(c, 401, "Failed to refresh token", err.Error())
        return
    }

    utils.SuccessMessage(c, 200, "Refresh success", res)
}