package model

import (
	"time"

	"github.com/guregu/null/v5"
)

type User struct {
	Username  	string
	Password  	string
	CreatedAt 	time.Time		`json:"createdAt"`
	UpdatedAt 	null.Time		`json:"updatedAt"`
	Balance  	[]*Balance		`gorm:"foreignKey:UserID"`
	ID        	uint			`gorm:"primaryKey"`
}

func (User) TableName() string {
	return "users"
}