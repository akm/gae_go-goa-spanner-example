// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "appengine": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/akm/gae_go-goa-spanner-example/design
// --out=$(GOPATH)/src/github.com/akm/gae_go-goa-spanner-example
// --version=v1.3.1

package app

import (
	"github.com/goadesign/goa"
)

// Book (default view)
//
// Identifier: application/vnd.book+json; view=default
type Book struct {
	Author string `form:"author" json:"author" yaml:"author" xml:"author"`
	Name   string `form:"name" json:"name" yaml:"name" xml:"name"`
	UserID int64  `form:"userId" json:"userId" yaml:"userId" xml:"userId"`
}

// Validate validates the Book media type instance.
func (mt *Book) Validate() (err error) {
	if mt.UserID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "userId"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Author == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "author"))
	}
	return
}

// User (default view)
//
// Identifier: application/vnd.user+json; view=default
type User struct {
	City   *string `form:"city,omitempty" json:"city,omitempty" yaml:"city,omitempty" xml:"city,omitempty"`
	Email  string  `form:"email" json:"email" yaml:"email" xml:"email"`
	Name   string  `form:"name" json:"name" yaml:"name" xml:"name"`
	UserID int64   `form:"userId" json:"userId" yaml:"userId" xml:"userId"`
}

// Validate validates the User media type instance.
func (mt *User) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	return
}

// UserCollection is the media type for an array of User (default view)
//
// Identifier: application/vnd.user+json; type=collection; view=default
type UserCollection []*User

// Validate validates the UserCollection media type instance.
func (mt UserCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}
