package handler

import (
	"errors"
	"math"
	"strconv"

	"github.com/AdnanRohmatKurniansah/tweet-go/internal/config"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/dto"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/repository"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/service"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentHandler struct {
	commentService service.CommentService
}

func NewCommentHandler(cfg *config.Config, db *gorm.DB) *CommentHandler {
	repo := repository.NewCommentRepository(db)
	svc := service.NewCommentService(repo)
	return &CommentHandler{commentService: svc}
}

func (h *CommentHandler) GetAll(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("postId"))
	if err != nil {
		utils.ErrorMessage(c, 400, "Invalid post id format", err.Error())
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	comments, total, err := h.commentService.GetComments(postId, page, limit)
	if err != nil {
		utils.ErrorMessage(c, 500, "Failed to fetch comments", err.Error())
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	utils.PaginationMessage(c, 200, "Comments fetched successfully", comments, total, page, limit, totalPages)
}

func (h *CommentHandler) GetById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorMessage(c, 400, "Invalid id format", err.Error())
		return
	}

	comment, err := h.commentService.GetCommentById(uint(id))
	if err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			utils.ErrorMessage(c, 404, "Comment not found", nil)
			return
		}

		utils.ErrorMessage(c, 500, "Failed to fetch comment", err.Error())
		return
	}

	utils.SuccessMessage(c, 200, "Comment fetched successfully", comment)
}

func (h *CommentHandler) Create(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		utils.ErrorMessage(c, 401, "Unauthorized", "User not authenticated")
		return
	}

	var req dto.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := utils.FormatValidationError(err, req)
		utils.ErrorMessage(c, 400, "Validation failed", validationErrors)
		return
	}

	comment, err := h.commentService.CreateComment(req, userId.(uint))
	if err != nil {
		if errors.Is(err, utils.ErrAlreadyExists) {
			utils.ErrorMessage(c, 409, "You have already commented on this post", nil)
			return
		}

		utils.ErrorMessage(c, 500, "Failed to create comment", err.Error())
		return
	}

	utils.SuccessMessage(c, 201, "Comment created successfully", comment)
}

func (h *CommentHandler) Update(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		utils.ErrorMessage(c, 401, "Unauthorized", "User not authenticated")
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorMessage(c, 400, "Invalid id format", err.Error())
		return
	}

	var req dto.UpdateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		validationError := utils.FormatValidationError(err, req)
		utils.ErrorMessage(c, 400, "Validation failed", validationError)
		return
	}

	comment, err := h.commentService.UpdateComment(uint(id), req, userId.(uint))
	if err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			utils.ErrorMessage(c, 404, "Comment not found", nil)
			return
		}

		if errors.Is(err, utils.ErrForbidden) {
			utils.ErrorMessage(c, 403, "Forbidden", nil)
			return
		}

		utils.ErrorMessage(c, 500, "Failed to update comment", err.Error())
		return
	}

	utils.SuccessMessage(c, 200, "Comment updated successfully", comment)
}

func (h *CommentHandler) Delete(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		utils.ErrorMessage(c, 401, "Unauthorized", "User not authenticated")
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorMessage(c, 400, "Invalid id format", err.Error())
		return
	}

	comment, err := h.commentService.DeleteComment(uint(id), userId.(uint))
	if err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			utils.ErrorMessage(c, 404, "Comment not found", nil)
			return
		}

		if errors.Is(err, utils.ErrForbidden) {
			utils.ErrorMessage(c, 403, "Forbidden", nil)
			return
		}

		utils.ErrorMessage(c, 500, "Failed to delete comment", err.Error())
		return
	}

	utils.SuccessMessage(c, 200, "Comment deleted successfully", comment)
}