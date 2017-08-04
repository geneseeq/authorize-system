// Package user contains the heart of the domain model.
package user

import (
	"errors"
	"strings"
	"time"

	"github.com/pborman/uuid"
)

// UserModel is user struct
type UserModel struct {
	ID           string
	Type         int
	Number       string
	Username     string
	Tele         string
	Gneder       bool
	Status       int
	Validity     bool
	Vip          bool
	Buildin      bool
	CreateUserID string
	CreateTime   time.Time
	Avatar       string
}

// NewUser is create instance
func NewUser(id string) *UserModel {
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

// TimeUtcToCst is format time
func TimeUtcToCst(t time.Time) time.Time {
	return t.Add(time.Hour * time.Duration(8))
}

// Repository is user interface
type Repository interface {
	Store(user *UserModel) error
	Find(id string) (*UserModel, error)
	FindAll() []*UserModel
	Remove(id string) error
	Update(id string, user *UserModel) error
}
