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
