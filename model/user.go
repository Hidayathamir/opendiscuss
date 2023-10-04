package model

import "time"

type User struct {
	ID        int       `gorm:"column:id"`
	Username  string    `gorm:"column:username"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}