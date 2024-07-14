package dto

import (
	"time"

	"github.com/guregu/null/v5"
	"goyave.dev/goyave/v5/util/typeutil"
)

// User the public user DTO. Used to show profiles for example.
type User struct {
	CreatedAt time.Time 				`json:"createdAt"`
	UpdatedAt null.Time 				`json:"updatedAt"`
	Username  string    				`json:"username"`
	ID        uint      				`json:"id"`
}

// InternalUser contains private user info that should not be exposed to clients.
type InternalUser struct {
	Password string						`json:"password"`
	User
}

type RegisterUser struct {
	Username string						`json:"username"`
	Password string						`json:"password" copier:"-"`
}

type UpdateUser struct {
	Username typeutil.Undefined[string]	`json:"username"`
	Password typeutil.Undefined[string]	`json:"password" copier:"-"`
}