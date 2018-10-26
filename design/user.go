package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var Int64Metadata = func() {
	Metadata("struct:field:type", "int64")
}

var UserPayload = Type("UserPayload", func() {
	// Member("user_id", String) // UUID
	Member("name", String)
	Member("email", String)
	Member("city", String)
	Required("name", "email", "city")
})

var User = MediaType("application/vnd.user+json", func() {
	Description("User")

	attrNames := []string{"user_id", "name", "email", "city"}
	Reference(UserPayload)
	Attributes(func() {
		for _, attrName := range attrNames {
			Attribute(attrName)
		}
		Required("user_id", "name", "email")
	})

	View("default", func() {
		for _, attrName := range attrNames {
			Attribute(attrName)
		}
	})
})

var _ = Resource("User", func() {
	BasePath("/users")
	DefaultMedia(User)

	Action("list", func() {
		Description("list")
		Routing(GET(""))
		Response(OK, CollectionOf(User))
		UseTrait(DefaultResponseTrait)
	})
	Action("create", func() {
		Description("create")
		Routing(POST(""))
		Payload(UserPayload)
		Response(Created, User)
		UseTrait(DefaultResponseTrait)
	})
	Action("show", func() {
		Description("show")
		Routing(GET("/:user_id"))
		Params(func() {
			Param("user_id", String) // UUID
		})
		Response(OK, User)
		UseTrait(DefaultResponseTrait)
	})
	Action("update", func() {
		Routing(PUT("/:user_id"))
		Params(func() {
			Param("user_id", String) // UUID
		})
		Payload(UserPayload)
		Response(OK, User)
		UseTrait(DefaultResponseTrait)
	})
	Action("delete", func() {
		Description("delete")
		Routing(DELETE("/:user_id"))
		Params(func() {
			Param("user_id", String) // UUID
		})
		Response(OK, User)
		UseTrait(DefaultResponseTrait)
	})
})
