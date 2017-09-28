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
	UnionID      string    `bson:"_id"` //唯一ID
	ID           string    `bson:"id"`
	Type         int       `bson:"type"`
	Number       string    `bson:"number"`
	Username     string    `bson:"username"`
	Tele         string    `bson:"tele"`
	Gneder       bool      `bson:"gender"`
	Status       int       `bson:"status"`
	Validity     bool      `bson:"validity"`
	Vip          bool      `bson:"vip"`
	Buildin      bool      `bson:"buildin"`
	CreateUserID string    `bson:"create_user_id"`
	CreateTime   time.Time `bson:"create_time"`
	UpdateTime   time.Time `bson:"update_time"`
	Avatar       string    `bson:"avatar"`
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
	ErrUnknown = errors.New("not find data.")
	ErrRemove  = errors.New("delete data is failed.")
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
