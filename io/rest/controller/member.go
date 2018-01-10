package controller

import (
	"github.com/aweisser/ocg-poc/io/rest"
	"github.com/aweisser/ocg-poc/io/rest/app"
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
	pi, _ := rest.GetMemberProfileInteraction(ctx)
	m, _ := pi.LoadByID(ctx.MemberID)

	res := &app.MemberSingle{
		Data: &app.Member{
			ID: &ctx.MemberID,
			Attributes: &app.MemberAttributes{
				Name: m.Name,
			},
		},
	}
	return ctx.OK(res)
	// MemberController_Show: end_implement
}
