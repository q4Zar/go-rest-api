package model

import (
	"time"

	"github.com/guregu/null/v5"
)

type User struct {
	Username  string
	Password  string
	Asset     []*Asset  `gorm:"foreignKey:UserID"`
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt null.Time `json:"updatedAt"`
}

func (User) TableName() string {
	return "users"
}
