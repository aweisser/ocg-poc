package controller

import (
	"github.com/aweisser/ocg-poc/cmd/ocg-rest-server/app"
	"github.com/aweisser/ocg-poc/domain/ocg"
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

	// TODO Call a usecase to get the member
	member := &ocg.Member{
		Name: "johndoe" + ctx.MemberID,
	}

	// Create response model
	res := &app.MemberSingle{
		Data: &app.Member{
			ID: &ctx.MemberID,
			Attributes: &app.MemberAttributes{
				Name: member.Name,
			},
		},
	}
	return ctx.OK(res)
	// MemberController_Show: end_implement
}
