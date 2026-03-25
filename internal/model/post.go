package model

import "time"

type Post struct {
    Id uint `gorm:"primaryKey" json:"id"`
    UserId uint `gorm:"not null;index" json:"user_id"`
    Title string `gorm:"type:varchar(255);not null" json:"title"`
    Content string `gorm:"not null" json:"content"`
    ImageUrl string `gorm:"type:text" json:"image_url"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
    
    User User `gorm:"foreignKey:UserId" json:"user"`
}