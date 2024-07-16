package dto

import (
	"time"

	"github.com/guregu/null/v5"
	"goyave.dev/goyave/v5/util/typeutil"
)

type Currency struct {
	ID        	uint		`json:"id"`
	Name		string		`json:"name"`
	CreatedAt	time.Time	`json:"createdAt"`
	UpdatedAt	null.Time	`json:"updatedAt"`
}

type CreateCurrency struct {
	Name    	string		`json:"name"`
}

type UpdateCurrency struct {
	Name    	typeutil.Undefined[string]	`json:"name"`
}
