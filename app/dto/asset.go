package dto

import (
	"goyave.dev/goyave/v5/util/typeutil"
)

type Asset struct {
	User *User `json:"owner,omitempty"`

	Balance   float64 `json:"balance,omitempty"`
	AssetType string  `json:"assetType,omitempty"`
	UserID    uint    `json:"userID,omitempty"`
	ID        uint    `json:"id,omitempty"`
}

type CreateAsset struct {
	Balance   float64 `json:"balance"`
	AssetType string  `json:"assetType"`
	UserID    uint    `json:"userID"`
}

type UpdateAsset struct {
	Balance typeutil.Undefined[float64] `json:"balance"`
}
