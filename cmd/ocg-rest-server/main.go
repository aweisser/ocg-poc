//go:generate goagen bootstrap -d github.com/aweisser/ocg-poc/cmd/ocg-rest-server/design

package main

import (
	"github.com/aweisser/ocg-poc/cmd/ocg-rest-server/app"
	"github.com/aweisser/ocg-poc/cmd/ocg-rest-server/controller"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("ocg")

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
