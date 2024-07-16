package model

import (
	"time"

	"github.com/guregu/null/v5"
	"gorm.io/gorm"
)

type		CurrencyName string

const (
    Euro	CurrencyName = "EUR"
    Dollar	CurrencyName = "USD"
)

type Currency struct {
	CreatedAt	time.Time
	UpdatedAt	null.Time
	DeletedAt	gorm.DeletedAt
	Name		CurrencyName`gorm:"type:enum('EUR', 'USD'),uniqueIndex:idx_nameowner"`
	ID			uint `gorm:"primarykey"`
}

func (Currency) TableName() string {
	return "currencies"
}