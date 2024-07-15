package user

import (
	vv "github.com/q4Zar/go-rest-api/http/validation"
	"goyave.dev/goyave/v5"
	v "goyave.dev/goyave/v5/validation"
)

func (ctrl *Controller) RegisterRequest(_ *goyave.Request) v.RuleSet {
	return v.RuleSet{
		{Path: v.CurrentElement, Rules: v.List{v.Required(), v.Object()}},
		{Path: "username", Rules: v.List{v.Required(), v.String(), v.Trim(), v.Between(3, 100)}},
		{Path: "password", Rules: v.List{v.Required(), v.String(), v.Between(6, 72), vv.Password()}},
	}
}

func (ctrl *Controller) UpdateRequest(_ *goyave.Request) v.RuleSet {
	return v.RuleSet{
		{Path: v.CurrentElement, Rules: v.List{v.Required(), v.Object()}},
		{Path: "username", Rules: v.List{v.String(), v.Trim(), v.Between(3, 100)}},
		{Path: "password", Rules: v.List{v.String(), v.Between(6, 72), vv.Password()}},
	}
}