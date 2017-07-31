// Package user contains the heart of the domain model.
package user

import (
	"errors"
	"strings"

	"github.com/pborman/uuid"
)

// TrackingID is not use
// type TrackingID string

// UserModel is user struct
type UserModel struct {
	ID       string
	Type     int
	Number   string
	Username string
	Tele     string
	Gneder   bool
	Status   int
	Validity bool
	Vip      bool
	Buildin  bool
}

// New is create instance
func New(id string) *UserModel {
	return &UserModel{
		ID: id,
	}
}

// NextTrackingID generates a new tracking ID.
// TODO: Move to infrastructure(?)
func NextTrackingID() string {
	return strings.Split(strings.ToUpper(uuid.New()), "-")[0]
}

// ErrUnknown is unkown user error
var (
	ErrUnknown = errors.New("unknown user")
)

// Repository is user interface
type Repository interface {
	Store(user *UserModel) error
	Find(id string) (*UserModel, error)
	FindAll() []*UserModel
	Remove(id string) error
}
