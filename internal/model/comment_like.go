package model

import "time"

type CommentLike struct {
	Id uint `gorm:"primaryKey" json:"id"`
	CommentId uint `gorm:"not null;index" json:"comment_id"`
	UserId uint `gorm:"not null;index" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Comment Comment `gorm:"foreignKey:CommentId" json:"comment"`
	User User `gorm:"foreignKey:UserId" json:"user"`
}