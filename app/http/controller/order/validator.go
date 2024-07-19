package order

import (
	"goyave.dev/goyave/v5"
	v "goyave.dev/goyave/v5/validation"
)

func (ctrl *Controller) CreateRequest(_ *goyave.Request) v.RuleSet {
	return v.RuleSet{
		{Path: v.CurrentElement, Rules: v.List{v.Required(), v.Object()}},
		{Path: "amount", Rules: v.List{v.Required(), v.Float64()}},
		{Path: "price", Rules: v.List{v.Required(), v.Float64()}},
		{Path: "order_pair", Rules: v.List{v.String(), v.Trim(), v.Between(3, 10)}},
		{Path: "side", Rules: v.List{v.String(), v.Trim(), v.Between(3, 10)}},
	}
}
