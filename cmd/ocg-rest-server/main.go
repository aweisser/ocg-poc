//go:generate goagen bootstrap -d github.com/aweisser/ocg-poc/cmd/ocg-rest-server/design

package main

import (
	"context"

	"github.com/aweisser/ocg-poc/io/dummy"
	"github.com/aweisser/ocg-poc/io/rest"
	"github.com/aweisser/ocg-poc/io/rest/app"
	"github.com/aweisser/ocg-poc/io/rest/controller"
	"github.com/aweisser/ocg-poc/usecase/member"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("ocg")

	// Create and inject business services through special context
	service.Context = newRestServerContext(service.Context)

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "member" controller
	c := controller.NewMemberController(service)
	app.MountMemberController(service, c)

	// Start service
	if err := service.ListenAndServe(":9000"); err != nil {
		service.LogError("startup", "err", err)
	}

}

// newRestServerContext creates a context for the OCG Rest Server where all necessary business services are injected.
func newRestServerContext(parentCtx context.Context) context.Context {
	return context.WithValue(parentCtx, rest.MemberProfileKey,
		&member.ProfileInteraction{
			MemberRepo: &dummy.Repo{},
		})
}
