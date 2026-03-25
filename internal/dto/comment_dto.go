package dto

import "time"

type CreateCommentRequest struct {
    PostId uint `json:"post_id" binding:"required"`
    Content string `json:"content" binding:"required,max=255"`
}

type UpdateCommentRequest struct {
    Content string `json:"content" binding:"max=255"`
}

type CommentResponse struct {
	Id uint `json:"id"`
	PostId uint `json:"post_id"`
	UserId uint `json:"user_id"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


