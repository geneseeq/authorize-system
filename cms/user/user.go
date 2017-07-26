// Package user contains the heart of the domain model.
package user

import (
	"errors"
	"strings"

	"github.com/pborman/uuid"
)

type TrackingID string

type UserModel struct {
	TrackingID TrackingID
	Type       int
	Number     string
	Username   string
	Gneder     bool
	Status     int
	Validity   bool
	Vip        bool
	Buildin    bool
}

func New(id TrackingID) *UserModel {
	return &UserModel{
		TrackingID: id,
	}
}

// NextTrackingID generates a new tracking ID.
// TODO: Move to infrastructure(?)
func NextTrackingID() TrackingID {
	return TrackingID(strings.Split(strings.ToUpper(uuid.New()), "-")[0])
}

var (
	ErrUnknown = errors.New("unknown user")
)

type Repository interface {
	Store(user *UserModel) error
	Find(id string) (*UserModel, error)
	FindAll() []*UserModel
}
