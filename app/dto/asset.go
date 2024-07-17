package dto

import (
	"goyave.dev/goyave/v5/util/typeutil"
)

type Asset struct {
	Amount     float64 `json:"amount,omitempty"`
	CurrencyID uint    `json:"currencyID,omitempty"`
	UserID     uint    `json:"authorID,omitempty"`
	ID         uint    `json:"id,omitempty"`
}

type CreateAsset struct {
	Amount     float64 `json:"amount"`
	CurrencyID uint    `json:"currencyID"`
	UserID     uint    `json:"userID"`
}

type UpdateAsset struct {
	Amount typeutil.Undefined[float64] `json:"amount"`
}
