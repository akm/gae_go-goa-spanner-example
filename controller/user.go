package controller

import (
	"github.com/akm/gae_go-goa-spanner-example/app"
	"github.com/goadesign/goa"
	"github.com/satori/go.uuid"

	"cloud.google.com/go/spanner"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
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

	appCtx := appengine.NewContext(ctx.Request)
	return createClient(appCtx, func(client *spanner.Client) error {
		u, err :=  uuid.NewV4()
		if err != nil {
			return err
		}
		userId := u.String()

		userColumns := []string{
			"UserId", "Name", "Email", "City",
		}
		m := []*spanner.Mutation{
			spanner.InsertOrUpdate("Users",
				userColumns,
				[]interface{}{
					userId,
					ctx.Payload.Name,
					ctx.Payload.Email,
					ctx.Payload.City,
				}),
		}
		if _, err := client.Apply(ctx, m); err != nil {
			log.Errorf(ctx, "Failed to insert\n")
			return err
		}
		return ctx.Created(&app.User{
			UserID: userId,
			Name: ctx.Payload.Name,
			Email: ctx.Payload.Email,
			City: &ctx.Payload.City,
		})
	})
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
