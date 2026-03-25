package handler

import (
	"strconv"

	"github.com/AdnanRohmatKurniansah/tweet-go/internal/config"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/repository"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/service"
	"github.com/AdnanRohmatKurniansah/tweet-go/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentLikeHandler struct {
	service service.CommentLikeService
}

func NewCommentLikeHandler(cfg *config.Config, db *gorm.DB) *CommentLikeHandler {
	repo := repository.NewCommentLikeRepository(db)
	svc := service.NewCommentLikeService(repo)

	return &CommentLikeHandler{svc}
}

func (h *CommentLikeHandler) LikeUnlike(c *gin.Context) {
	userId, _ := c.Get("userId")

	id, err := strconv.ParseUint(c.Param("commentId"), 10, 32)
	if err != nil {
		utils.ErrorMessage(c, 400, "Invalid Id format", err.Error())
		return
	}

	liked, total, err := h.service.LikeOrUnlike(uint(id), userId.(uint))
	if err != nil {
		utils.ErrorMessage(c, 500, "Failed toggle like", err.Error())
		return
	}

	utils.SuccessMessage(c, 200, "Success", gin.H{
		"liked": liked,
		"total_likes": total,
	})
}

func (h *CommentLikeHandler) GetTotalLikes(c *gin.Context) {
	commentId, err := strconv.ParseUint(c.Param("commentId"), 10, 32)
	if err != nil {
		utils.ErrorMessage(c, 400, "Invalid comment id", err.Error())
		return
	}

	totalLikes, err := h.service.GetTotalLikes(uint(commentId))
	if err != nil {
		utils.ErrorMessage(c, 500, "Failed to get total likes", err.Error())
		return
	}

	utils.SuccessMessage(c, 200, "Total likes fetched successfully", gin.H{
		"comment_id": commentId,
		"total_likes": totalLikes,
	})
}