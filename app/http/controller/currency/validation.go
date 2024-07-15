package currency

import (
	"goyave.dev/goyave/v5"
	v "goyave.dev/goyave/v5/validation"
)

func (ctrl *Controller) CreateRequest(_ *goyave.Request) v.RuleSet {
	return v.RuleSet{
		{Path: v.CurrentElement, Rules: v.List{v.Required(), v.Object()}},
		{Path: "name", Rules: v.List{v.Required(), v.String(), v.Trim(), v.Between(1, 200)}},
		{Path: "amount", Rules: v.List{v.Required(), v.Int()}},
	}
}

func (ctrl *Controller) UpdateRequest(_ *goyave.Request) v.RuleSet {
	return v.RuleSet{
		{Path: v.CurrentElement, Rules: v.List{v.Required(), v.Object()}},
		{Path: "name", Rules: v.List{v.String(), v.Trim(), v.Between(1, 200)}},
		{Path: "amount", Rules: v.List{v.Int()}},
	}
}