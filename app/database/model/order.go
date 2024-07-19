package model

import (
	"time"

	"github.com/guregu/null/v5"
	"gorm.io/gorm"
)

// type OrderSideName string

// const (
// 	Buy  OrderSideName = "BUY"
// 	Sell OrderSideName = "SELL"
// )

// type OrderPairName string

// const (
// 	EuroDollar OrderSideName = "EUR-USD"
// 	DollarEuro OrderSideName = "USD-EUR"
// )

// type OrderStatusName string

// const (
// 	Pending OrderSideName = "Pending"
// 	Filled  OrderSideName = "Filled"
// )

type Order struct {
	User *User

	Side      string  `gorm:"type:enum('BUY', 'SELL'), required"`
	Amount    float64 `gorm:"required,gt=0"`
	Price     float64 `gorm:"required,gt=0"`
	AssetPair string  `gorm:"required"`
	Status    string  `gorm:"default:'Pending'"`
	UserID    uint
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt null.Time
	DeletedAt gorm.DeletedAt
}

func (Order) TableName() string {
	return "orders"
}
