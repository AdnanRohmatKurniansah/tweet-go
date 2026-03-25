package model

import "time"

type PostLike struct {
	Id uint `gorm:"primaryKey" json:"id"`
	PostId uint `gorm:"not null;index" json:"post_id"`
	UserId uint `gorm:"not null;index" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Post Post `gorm:"foreignKey:PostId" json:"post"`
	User User `gorm:"foreignKey:UserId" json:"user"`
}