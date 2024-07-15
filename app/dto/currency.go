package dto

import (
	"time"

	"github.com/guregu/null/v5"
	"goyave.dev/goyave/v5/util/typeutil"
)

type Currency struct {
	Owner		*User `json:"owner,omitempty"`

	CreatedAt	time.Time	`json:"createdAt"`
	UpdatedAt	null.Time	`json:"updatedAt"`
	Name		string		`json:"name"`
	Amount		uint		`json:"amount"`
	OwnerID 	uint		`json:"ownerID"`
	ID        	uint		`json:"id"`
}

type CreateCurrency struct {
	Name    	string		`json:"name"`
	Amount		uint		`json:"amount"`
	OwnerID 	uint		`json:"ownerID"`
}

type UpdateCurrency struct {
	Name    	typeutil.Undefined[string]	`json:"name"`
	Amount		typeutil.Undefined[uint]	`json:"amount"`
}

type ShowCurrency struct {
	Name    	string		`json:"name"`
	Amount		uint		`json:"amount"`
	OwnerID 	uint		`json:"ownerID"`
}