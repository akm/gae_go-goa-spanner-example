package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var Int64Metadata = func() {
	Metadata("struct:field:type", "int64")
}

var UserPayload = Type("UserPayload", func() {
	Member("userId", Integer, Int64Metadata)
	Member("name", String)
	Member("email", String)
	Member("city", String)
	Required("userId", "name", "email", "city")
})

var User = MediaType("application/vnd.user+json", func() {
	Description("User")

	attrNames := []string{"userId", "name", "email", "city"}
	Reference(UserPayload)
	Attributes(func() {
		for _, attrName := range attrNames {
			Attribute(attrName)
		}
		Required("userId", "name", "email")
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
		Routing(GET("/:userId"))
		Params(func() {
			Param("userId", Integer, Int64Metadata)
		})
		Response(OK, User)
		UseTrait(DefaultResponseTrait)
	})
	Action("update", func() {
		Routing(PUT("/:userId"))
		Params(func() {
			Param("userId", Integer, Int64Metadata)
		})
		Payload(UserPayload)
		Response(OK, User)
		UseTrait(DefaultResponseTrait)
	})
	Action("delete", func() {
		Description("delete")
		Routing(DELETE("/:userId"))
		Params(func() {
			Param("userId", Integer, Int64Metadata)
		})
		Response(OK, User)
		UseTrait(DefaultResponseTrait)
	})
})
