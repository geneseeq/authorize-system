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
	GetAllUser() ([]User, error)
	PutUser(id string, user User) error
	DeleteUser(id string) error
}

// User is a user base info
type User struct {
	ID             string `json:"id"`
	Type           int    `json:"type,omitempty"` //"type":"医生/教师/个人/员工/企业"
	Number         string `json:"number,omitempty"`
	Username       string `json:"username,omitempty"`
	Tele           string `json:"tele,omitempty"`
	Gneder         bool   `json:"gender,omitempty"`
	Status         int    `json:"status,omitempty"`
	Validity       bool   `json:"validity,omitempty"`
	Vip            bool   `json:"vip,omitempty"`
	Buildin        bool   `json:"buildin,omitempty"`
	Create_user_id string `json:"create_user_id,omitempty"`
	Create_time    string `json:"create_time,omitempty"`
	Avatar         string `json:"avatar,omitempty"`
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

func (s *service) GetAllUser() ([]User, error) {
	var result []User
	for _, c := range s.users.FindAll() {
		result = append(result, usermodelToUser(c))
	}
	return result, nil
}

func (s *service) PutUser(id string, user User) error {
	_, err := s.GetUser(id)
	if err != nil {
		return ErrInconsistentIDs
	}
	err = s.users.Update(id, userToUsermodel(user))
	return err
}

func (s *service) DeleteUser(id string) error {
	if id == "" {
		return ErrInvalidArgument
	}
	error := s.users.Remove(id)
	if error != nil {
		return ErrNotFound
	}
	return nil
}
func userToUsermodel(u User) *user.UserModel {

	return &user.UserModel{
		ID:             u.ID,
		Type:           u.Type,
		Number:         u.Number,
		Username:       u.Username,
		Gneder:         u.Gneder,
		Tele:           u.Tele,
		Status:         u.Status,
		Validity:       u.Validity,
		Vip:            u.Vip,
		Buildin:        u.Buildin,
		Create_user_id: u.Create_user_id,
		Create_time:    u.Create_time,
		Avatar:         u.Avatar,
	}
}

func usermodelToUser(c *user.UserModel) User {
	return User{
		ID:             c.ID,
		Type:           c.Type,
		Number:         c.Number,
		Username:       c.Username,
		Gneder:         c.Gneder,
		Tele:           c.Tele,
		Status:         c.Status,
		Validity:       c.Validity,
		Vip:            c.Vip,
		Buildin:        c.Buildin,
		Create_user_id: c.Create_user_id,
		Create_time:    c.Create_time,
		Avatar:         c.Avatar,
	}
}
