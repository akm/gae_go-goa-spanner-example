package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var BookPayload = Type("BookPayload", func() {
	Member("userId", Integer, func() {
		Metadata("struct:field:type", "int64")
	})
	Member("name", String)
	Member("author", String)
	Required("userId", "name", "author")
})

var Book = MediaType("application/vnd.book+json", func() {
	Description("Book")

	attrNames := []string{"userId", "name", "author"}
	Reference(BookPayload)
	Attributes(func() {
		for _, attrName := range attrNames {
			Attribute(attrName)
		}
		Required(attrNames...)
	})

	View("default", func() {
		for _, attrName := range attrNames {
			Attribute(attrName)
		}
	})
})

var _ = Resource("Book", func() {
	BasePath("/books")
	DefaultMedia(Book)

	Action("list", func() {
		Description("list")
		Routing(GET(""))
		Response(OK, CollectionOf(Book))
		UseTrait(DefaultResponseTrait)
	})
	Action("create", func() {
		Description("create")
		Routing(POST(""))
		Payload(BookPayload)
		Response(Created, Book)
		UseTrait(DefaultResponseTrait)
	})
	Action("show", func() {
		Description("show")
		Routing(GET("/:name"))
		Params(func() {
			Param("name", String)
		})
		Response(OK, Book)
		UseTrait(DefaultResponseTrait)
	})
	Action("update", func() {
		Routing(PUT("/:name"))
		Params(func() {
			Param("name", String)
		})
		Payload(BookPayload)
		Response(OK, Book)
		UseTrait(DefaultResponseTrait)
	})
	Action("delete", func() {
		Description("delete")
		Routing(DELETE("/:name"))
		Params(func() {
			Param("name", String)
		})
		Response(OK, Book)
		UseTrait(DefaultResponseTrait)
	})
})
