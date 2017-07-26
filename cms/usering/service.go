// Package booking provides the use-case of booking a cargo. Used by views
// facing an administrator.
package booking

import (
	"errors"
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
	// PostUser(user User) error
	// GetAllUser() ([]User, error)
	// PutUser(id user.TrackingID, user User) error
	// DeleteUser(id user.TrackingID) error
}

// User is a user base info
type User struct {
	ID       string `json:"id"`
	Type     int    `json:"type"` //"type":"医生/教师/个人/员工/企业"
	Number   string `json:"number"`
	Username string `json:"username"`
	Gneder   bool   `json:"gender"`
	Status   int    `json:"status"`
	Validity bool   `json:"validity"`
	Vip      bool   `json:"vip"`
	Buildin  bool   `json:"buildin"`
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

// func (s *service) PostUser(user User) error {
// 	uid := user.NextTrackingID()
// 	u := user.New(user.ID)
// 	return s.users.Store(u)
// }

func (s *service) GetUser(id string) error {
	return s.users.Find(id)
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
