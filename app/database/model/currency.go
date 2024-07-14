package model

import (
	"time"

	"github.com/guregu/null/v5"
	"gorm.io/gorm"
)

type Currency struct {
	Owner		*User
	CreatedAt	time.Time
	UpdatedAt	null.Time
	DeletedAt	gorm.DeletedAt
	Name		string
	OwnerID		uint
	ID			uint `gorm:"primarykey"`
}

func (Currency) TableName() string {
	return "currencies"
}