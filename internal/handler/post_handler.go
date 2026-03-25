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

type PostHandler struct {
	postService service.PostService
}

func NewPostHandler(cfg *config.Config, db *gorm.DB) *PostHandler {
	repo := repository.NewPostRepository(db)
	postSvc := service.NewPostService(repo)
	return &PostHandler{postService: postSvc}
}

func (h *PostHandler) GetAll(c *gin.Context) {
	page, _  := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	posts, total, err := h.postService.GetPosts(page, limit)
	if err != nil {
		utils.ErrorMessage(c, 500, "Failed to fetch posts", err.Error())
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	utils.PaginationMessage(c, 200, "Posts fetched successfully", posts, total, page, limit, totalPages)
}

func (h *PostHandler) GetById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorMessage(c, 400, "Invalid Id format", err.Error())
		return
	}

	post, err := h.postService.GetPostById(uint(id))
	if err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			utils.ErrorMessage(c, 404, "Post not found", nil)
			return
		}

		utils.ErrorMessage(c, 500, "Failed to fetch post", err.Error())
		return
	}

	utils.SuccessMessage(c, 200, "Post fetched successfully", post)
}

func (h *PostHandler) Create(c *gin.Context) {
	var req dto.CreatePostRequest
	if err := c.ShouldBind(&req); err != nil {
		validationErrors := utils.FormatValidationError(err, req)

		utils.ErrorMessage(c, 400, "Validation failed", validationErrors)
		return
	}

	userId, exists := c.Get("userId")
	if !exists {
		utils.ErrorMessage(c, 401, "Unauthorized", "User not authenticated")
		return
	}

	imageUrl, err := utils.SaveImage(c, "image_url", "posts", true)
	if err != nil {
		utils.ErrorMessage(c, 400, "Invalid image", err.Error())
		return
	}

	post, err := h.postService.CreatePost(req, userId.(uint), imageUrl)
	if err != nil {
		utils.ErrorMessage(c, 500, "Failed to create post", err.Error())
		return
	}

	utils.SuccessMessage(c, 201, "Post created successfully", post)
}

func (h *PostHandler) Update(c *gin.Context) {
	_, exists := c.Get("userId")
	if !exists {
		utils.ErrorMessage(c, 401, "Unauthorized", "User not authenticated")
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorMessage(c, 400, "Invalid Id format", err.Error())
		return
	}

	var req dto.UpdatePostRequest
	if err := c.ShouldBind(&req); err != nil {
		validationError := utils.FormatValidationError(err, req)

		utils.ErrorMessage(c, 400, "Validation failed", validationError)
		return
	}

	imageUrl, err := utils.SaveImage(c, "image_url", "posts", false)
	if err != nil {
		utils.ErrorMessage(c, 400, "Invalid image", err.Error())
		return
	}

	post, err := h.postService.UpdatePost(uint(id), req, imageUrl)
	if err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			utils.ErrorMessage(c, 404, "Post not found", nil)
			return
		}
		utils.ErrorMessage(c, 500, "Failed to update post", err.Error())
		return
	}

	utils.SuccessMessage(c, 200, "Post updated successfully", post)
}

func (h *PostHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorMessage(c, 400, "Invalid Id format", err.Error())
		return
	}

	post, err := h.postService.DeletePost(uint(id))
	if ; err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			utils.ErrorMessage(c, 404, "Post not found", nil)
			return
		}

		utils.ErrorMessage(c, 500, "Failed to delete post", err.Error())
		return
	}

	utils.SuccessMessage(c, 200, "Post deleted successfully", post)
}