package controller

import (
	"github.com/aweisser/ocg-poc/cmd/ocg-rest-server/app"
	"github.com/goadesign/goa"
)

// MemberController implements the member resource.
type MemberController struct {
	*goa.Controller
}

// NewMemberController creates a member controller.
func NewMemberController(service *goa.Service) *MemberController {
	return &MemberController{Controller: service.NewController("MemberController")}
}

// Show runs the show action.
func (c *MemberController) Show(ctx *app.ShowMemberContext) error {
	// MemberController_Show: start_implement

	// Put your logic here

	res := &app.MemberSingle{
		Data: &app.Member{
			ID: &ctx.MemberID,
			Attributes: &app.MemberAttributes{
				Name: "johndoe" + ctx.MemberID,
			},
		},
	}
	return ctx.OK(res)
	// MemberController_Show: end_implement
}
