package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null"` // 密碼不會在 JSON 中返回
	Email     string    `json:"email" gorm:"unique;not null"`
	Balance   int       `json:"balance" gorm:"default:0"` // 用戶餘額（虛擬貨幣）
	BirthDate time.Time `json:"birth_date" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
