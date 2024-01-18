package model

import "gorm.io/gorm"

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Passwd    string `json:"passwd"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`

	gorm.Model
}

func (User) TableName() string {
	return "users"
}
