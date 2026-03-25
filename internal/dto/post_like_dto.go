package dto

type PostLikeRequest struct {
	PostId uint `json:"post_id" binding:"required"`
}