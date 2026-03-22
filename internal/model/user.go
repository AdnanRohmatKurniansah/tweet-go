package model

type User struct {
    Id uint `gorm:"primaryKey" json:"id"`
    Name string `gorm:"size:255" json:"name"`
    Email string `gorm:"uniqueIndex;not null;size:100" json:"email"`
    Phone string `gorm:"size:30" json:"phone"`
    Password string `gorm:"not null" json:"-"`
}