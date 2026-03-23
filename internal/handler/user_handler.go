package handler

import (
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/config"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/dto"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/repository"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/service"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(cfg *config.Config, db *gorm.DB) *UserHandler {
	repo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(cfg, repo)

	return &UserHandler{
		userService: userSvc,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
    var req dto.RegisterRequest
    if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := utils.FormatValidationError(err, req)

        utils.ErrorMessage(c, 400, "Validation failed", validationErrors)
		return
	}

    res, err := h.userService.Register(req); 
    
    if err != nil {
        utils.ErrorMessage(c, 500, "Failed to register user", err.Error())
        return
    }

    utils.SuccessMessage(c, 201, "Register success", res)
}

func (h *UserHandler) Login(c *gin.Context) {
    var req dto.LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        validationErrors := utils.FormatValidationError(err, req)

        utils.ErrorMessage(c, 400, "Validation failed", validationErrors)
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
        validationErrors := utils.FormatValidationError(err, req)

        utils.ErrorMessage(c, 400, "Validation failed", validationErrors)
        return
    }

    res, err := h.userService.Refresh(req)
    if err != nil {
        utils.ErrorMessage(c, 401, "Failed to refresh token", err.Error())
        return
    }

    utils.SuccessMessage(c, 200, "Refresh success", res)
}