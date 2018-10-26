// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "appengine": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=github.com/akm/gae_go-goa-spanner-example/design
// --out=$(GOPATH)/src/github.com/akm/gae_go-goa-spanner-example
// --version=v1.3.1

package app

import (
	"fmt"
	"strings"
)

// BookHref returns the resource href.
func BookHref(name interface{}) string {
	paramname := strings.TrimLeftFunc(fmt.Sprintf("%v", name), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/books/%v", paramname)
}

// UserHref returns the resource href.
func UserHref(userID interface{}) string {
	paramuserID := strings.TrimLeftFunc(fmt.Sprintf("%v", userID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/users/%v", paramuserID)
}
