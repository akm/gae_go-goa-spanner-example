// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "appengine": Application User Types
//
// Command:
// $ goagen
// --design=github.com/akm/gae_go-goa-spanner-example/design
// --out=$(GOPATH)/src/github.com/akm/gae_go-goa-spanner-example
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
)

// bookPayload user type.
type bookPayload struct {
	Author *string `form:"author,omitempty" json:"author,omitempty" yaml:"author,omitempty" xml:"author,omitempty"`
	Name   *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	UserID *string `form:"user_id,omitempty" json:"user_id,omitempty" yaml:"user_id,omitempty" xml:"user_id,omitempty"`
}

// Validate validates the bookPayload type instance.
func (ut *bookPayload) Validate() (err error) {
	if ut.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "user_id"))
	}
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	if ut.Author == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "author"))
	}
	return
}

// Publicize creates BookPayload from bookPayload
func (ut *bookPayload) Publicize() *BookPayload {
	var pub BookPayload
	if ut.Author != nil {
		pub.Author = *ut.Author
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.UserID != nil {
		pub.UserID = *ut.UserID
	}
	return &pub
}

// BookPayload user type.
type BookPayload struct {
	Author string `form:"author" json:"author" yaml:"author" xml:"author"`
	Name   string `form:"name" json:"name" yaml:"name" xml:"name"`
	UserID string `form:"user_id" json:"user_id" yaml:"user_id" xml:"user_id"`
}

// Validate validates the BookPayload type instance.
func (ut *BookPayload) Validate() (err error) {
	if ut.UserID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "user_id"))
	}
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}
	if ut.Author == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "author"))
	}
	return
}

// userPayload user type.
type userPayload struct {
	City  *string `form:"city,omitempty" json:"city,omitempty" yaml:"city,omitempty" xml:"city,omitempty"`
	Email *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	Name  *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the userPayload type instance.
func (ut *userPayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "email"))
	}
	if ut.City == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "city"))
	}
	return
}

// Publicize creates UserPayload from userPayload
func (ut *userPayload) Publicize() *UserPayload {
	var pub UserPayload
	if ut.City != nil {
		pub.City = *ut.City
	}
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	return &pub
}

// UserPayload user type.
type UserPayload struct {
	City  string `form:"city" json:"city" yaml:"city" xml:"city"`
	Email string `form:"email" json:"email" yaml:"email" xml:"email"`
	Name  string `form:"name" json:"name" yaml:"name" xml:"name"`
}

// Validate validates the UserPayload type instance.
func (ut *UserPayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "email"))
	}
	if ut.City == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "city"))
	}
	return
}
