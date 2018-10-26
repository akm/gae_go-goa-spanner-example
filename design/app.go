package design

import (
	"os"

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

const DefaultResponseTrait = "DefaultResponseTrait"

func env(key, defaultValue string) string {
	r := os.Getenv(key)
	if r == "" {
		return defaultValue
	}
	return r
}

var _ = API("appengine", func() {
	Title("The appengine example")
	Description("A simple appengine example")

	Host(env("GOA_HOST", "localhost:8080"))
	Scheme(env("GOA_SCHEME", "http"))

	BasePath("/")
	Origin("*", func() {
		Methods("GET", "POST", "PUT", "DELETE", "OPTIONS")
		MaxAge(600)
		Credentials()
	})

	Trait(DefaultResponseTrait, func() {
		Response(Unauthorized, ErrorMedia)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
		Response(Conflict, ErrorMedia)
	})
})
