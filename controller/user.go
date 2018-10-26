package controller

import (
	"github.com/akm/gae_go-goa-spanner-example/app"
	"github.com/goadesign/goa"
)

// UserController implements the User resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a User controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	// UserController_Create: start_implement

	// Put your logic here

	return nil
	// UserController_Create: end_implement
}

// Delete runs the delete action.
func (c *UserController) Delete(ctx *app.DeleteUserContext) error {
	// UserController_Delete: start_implement

	// Put your logic here

	res := &app.User{}
	return ctx.OK(res)
	// UserController_Delete: end_implement
}

// List runs the list action.
func (c *UserController) List(ctx *app.ListUserContext) error {
	// UserController_List: start_implement

	// Put your logic here

	res := app.UserCollection{}
	return ctx.OK(res)
	// UserController_List: end_implement
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	// UserController_Show: start_implement

	// Put your logic here

	res := &app.User{}
	return ctx.OK(res)
	// UserController_Show: end_implement
}

// Update runs the update action.
func (c *UserController) Update(ctx *app.UpdateUserContext) error {
	// UserController_Update: start_implement

	// Put your logic here

	res := &app.User{}
	return ctx.OK(res)
	// UserController_Update: end_implement
}
