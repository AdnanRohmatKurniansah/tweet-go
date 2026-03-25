package model

import "time"

type Comment struct {
    Id uint `gorm:"primaryKey" json:"id"`
    PostId uint `gorm:"not null;index" json:"post_id"`
    UserId uint `gorm:"not null;index" json:"user_id"`
    Content string `gorm:"not null" json:"content"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
    
	Post User `gorm:"foreignKey:PostId" json:"post"`
    User User `gorm:"foreignKey:UserId" json:"user"`
}