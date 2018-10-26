//go:generate goagen bootstrap -d github.com/akm/gae_go-goa-spanner-example/design

package main

import (
	"github.com/akm/gae_go-goa-spanner-example/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("appengine")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "User" controller
	c := NewUserController(service)
	app.MountUserController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
