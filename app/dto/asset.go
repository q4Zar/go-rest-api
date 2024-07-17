package dto

import (
	"time"

	"github.com/guregu/null/v5"
	"goyave.dev/goyave/v5/util/typeutil"
)

type Asset struct {
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  null.Time `json:"updatedAt"`
	Amount     float64   `json:"amount"`
	CurrencyID uint      `json:"currencyID"`
	UserID     uint      `json:"authorID"`
	ID         uint      `json:"id"`
}

type IndexAsset struct {
	Amount     float64 `json:"amount"`
	CurrencyID uint    `json:"currencyID"`
}

type CreateAsset struct {
	Amount     float64 `json:"amount"`
	CurrencyID uint    `json:"currencyID"`
	UserID     uint    `json:"userID"`
}

type UpdateAsset struct {
	Amount typeutil.Undefined[float64] `json:"amount"`
}
