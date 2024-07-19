package dto

type Order struct {
	User *User `json:"owner,omitempty"`

	Amount    float64 `json:"amount,omitempty"`
	Price     float64 `json:"price,omitempty"`
	AssetPair string  `json:"assetPair,omitempty"`
	Side      string  `json:"side,omitempty"`
	Status    string  `json:"status,omitempty"`
	UserID    uint    `json:"userID,omitempty"`
	ID        uint    `json:"id,omitempty"`
}

type CreateOrder struct {
	Amount    float64 `json:"amount"`
	Price     float64 `json:"price"`
	AssetPair string  `json:"assetPair"`
	Side      string  `json:"side"`
	UserID    uint    `json:"userID"`
}
