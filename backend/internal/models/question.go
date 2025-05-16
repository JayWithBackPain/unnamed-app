package models

import (
	"time"
)

type Question struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	UserID     uint      `json:"user_id" gorm:"not null"`
	Content    string    `json:"content" gorm:"not null"`
	Answer     string    `json:"answer"`
	Status     string    `json:"status" gorm:"default:'pending'"` // pending, answered, failed
	Cost       int       `json:"cost" gorm:"not null"`            // 問題消耗的虛擬貨幣
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	AnsweredAt time.Time `json:"answered_at"`
	User       User      `json:"user" gorm:"foreignKey:UserID"`
}
