package dto

import (
	"goyave.dev/goyave/v5/util/typeutil"
)

type Asset struct {
	Balance   float64 `json:"balance"`
	AssetType string  `json:"assetType"`
	UserID    uint    `json:"userID"`
	ID        uint    `json:"id"`
}

type CreateAsset struct {
	Balance   float64 `json:"balance"`
	AssetType string  `json:"assetType"`
	UserID    uint    `json:"userID"`
}

type UpdateAsset struct {
	Balance typeutil.Undefined[float64] `json:"balance"`
}
