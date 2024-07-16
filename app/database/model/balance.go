package model

import (
	"time"

	"github.com/guregu/null/v5"
	"gorm.io/gorm"
)

type Balance struct {
	CreatedAt	time.Time
	UpdatedAt	null.Time
	DeletedAt	gorm.DeletedAt
	Amount		float64
	CurrencyID	uint `gorm:"uniqueIndex:idx_unique_currencyid_&_userid"`
	UserID		uint `gorm:"uniqueIndex:idx_unique_currencyid_&_userid"`
	ID			uint `gorm:"primarykey"`
}

func (Balance) TableName() string {
	return "balances"
}
