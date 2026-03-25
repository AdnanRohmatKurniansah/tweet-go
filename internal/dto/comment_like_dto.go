package dto

type CommentLikeRequest struct {
	CommentId uint `json:"comment_id" binding:"required"`
}