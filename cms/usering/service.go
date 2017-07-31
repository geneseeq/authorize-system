// Package usering provides the use-case of booking a cargo. Used by views
// facing an administrator.
package usering

import (
	"errors"

	"github.com/geneseeq/authorize-system/cms/user"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var (
	ErrInvalidArgument = errors.New("invalid argument")
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

// Service is the interface that provides booking methods.
type Service interface {
	GetUser(id string) (User, error)
	PostUser(user User) error
	// GetAllUser() ([]User, error)
	// PutUser(id user.TrackingID, user User) error
	// DeleteUser(id user.TrackingID) error
}

// User is a user base info
type User struct {
	ID       string `json:"id"`
	Type     int    `json:"type,omitempty"` //"type":"医生/教师/个人/员工/企业"
	Number   string `json:"number,omitempty"`
	Username string `json:"username,omitempty"`
	Tele     string `json:"telephone,omitempty"`
	Gneder   bool   `json:"gender,omitempty"`
	Status   int    `json:"status,omitempty"`
	Validity bool   `json:"validity,omitempty"`
	Vip      bool   `json:"vip,omitempty"`
	Buildin  bool   `json:"buildin,omitempty"`
}

type service struct {
	users user.Repository
}

// NewService creates a booking service with necessary dependencies.
func NewService(users user.Repository) Service {
	return &service{
		users: users,
	}
}

func (s *service) PostUser(u User) error {
	return s.users.Store(userToUsermodel(u))
}

func (s *service) GetUser(id string) (User, error) {
	if id == "" {
		return User{}, ErrInvalidArgument
	}
	c, error := s.users.Find(id)
	if error != nil {
		return User{}, ErrNotFound
	}
	return usermodelToUser(c), nil
}

// func (s *service) GetAllUser() ([]User, error) {
// 	uid := user.NextTrackingID()
// 	u := user.New(id)
// 	return s.users.Store(u)
// }

// func (s *service) PutUser(user User) error {
// 	uid := user.NextTrackingID()
// 	u := user.New(id)
// 	return s.users.Store(u)
// }

// func (s *service) DeleteUser(user User) error {
// 	uid := user.NextTrackingID()
// 	u := user.New(id)
// 	return s.users.Store(u)
// }
func userToUsermodel(u User) *user.UserModel {

	return &user.UserModel{
		ID:       u.ID,
		Type:     u.Type,
		Number:   u.Number,
		Username: u.Username,
		Gneder:   u.Gneder,
		Status:   u.Status,
		Validity: u.Validity,
		Vip:      u.Vip,
		Buildin:  u.Buildin,
	}
}

func usermodelToUser(c *user.UserModel) User {
	return User{
		ID:       c.ID,
		Type:     c.Type,
		Number:   c.Number,
		Username: c.Username,
		Gneder:   c.Gneder,
		Status:   c.Status,
		Validity: c.Validity,
		Vip:      c.Vip,
		Buildin:  c.Buildin,
	}
}
