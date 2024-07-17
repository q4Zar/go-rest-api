package model

import (
	"time"

	"github.com/guregu/null/v5"
	"gorm.io/gorm"
)

type CurrencyName string

const (
	Euro   CurrencyName = "EUR"
	Dollar CurrencyName = "USD"
)

type Asset struct {
	User *User

	CreatedAt time.Time
	UpdatedAt null.Time
	DeletedAt gorm.DeletedAt
	Balance   float64
	AssetType CurrencyName `gorm:"type:enum('EUR', 'USD'),uniqueIndex:idx_unique_assettype_&_userid"`
	UserID    uint         `gorm:"uniqueIndex:idx_unique_assettype_&_userid"`
	ID        uint         `gorm:"primarykey"`
}

func (Asset) TableName() string {
	return "assets"
}
