package model

import "time"

type User struct {
	ID        int       `gorm:"column:id"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (User) Tablename() string {
	return "users"
}

const (
	USER_TABLE_NAME = "users"
	USER_ID         = "id"
	USER_USERNAME   = "username"
	USER_PASSWORD   = "password"
	USER_CREATED_AT = "created_at"
	USER_UPDATED_AT = "updated_at"
)
