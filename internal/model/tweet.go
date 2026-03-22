package model

type Tweet struct {
    Id uint `gorm:"primaryKey" json:"id"`
    Content string `gorm:"type:text;not null" json:"content"`
    UserId uint `gorm:"not null" json:"userId"`
    User User `gorm:"foreignKey:UserId" json:"user"`
}