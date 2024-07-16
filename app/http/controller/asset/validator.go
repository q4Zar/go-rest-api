package asset

import (
	"goyave.dev/goyave/v5"
	v "goyave.dev/goyave/v5/validation"
)

func (ctrl *Controller) CreateRequest(_ *goyave.Request) v.RuleSet {
	return v.RuleSet{
		{Path: v.CurrentElement, Rules: v.List{v.Required(), v.Object()}},
		{Path: "amount", Rules: v.List{v.Required(), v.Float64()}},
	}
}

func (ctrl *Controller) UpdateRequest(_ *goyave.Request) v.RuleSet {
	return v.RuleSet{
		{Path: v.CurrentElement, Rules: v.List{v.Required(), v.Object()}},
		{Path: "amount", Rules: v.List{v.Required(), v.Float64()}},
	}
}

// func (ctrl *Controller) GetRequest(_ *goyave.Request) v.RuleSet {
// 	return v.RuleSet{
// 		{Path: v.CurrentElement, Rules: v.List{v.Required(), v.Object()}},
// 		{Path: "name", Rules: v.List{v.Int()}},
// 	}
// }