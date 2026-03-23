package dto

import "time"

type CreatePostRequest struct {
	Title string `form:"title" binding:"required,max=255"`
	Content string `form:"content" binding:"required"`
}

type UpdatePostRequest struct {
	Title string `form:"title"`
	Content string `form:"content"`
}

type PostResponse struct {
	Id uint `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	ImageUrl string `json:"image_url"`
	UserId uint `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


