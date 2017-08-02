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
	PostUser(user []User) ([]string, error)
	GetAllUser() ([]User, error)
	PutUser(id string, user User) error
	DeleteUser(id string) error
	DeleteMultiUser(listid []string) ([]string, error)
	PutMultiUser(user []User) ([]string, error)
}

// User is a user base info
type User struct {
	ID             string `json:"id"`
	Type           int    `json:"type"` //"type":"医生/教师/个人/员工/企业"
	Number         string `json:"number"`
	Username       string `json:"username"`
	Tele           string `json:"tele"`
	Gneder         bool   `json:"gender"`
	Status         int    `json:"status"`
	Validity       bool   `json:"validity"`
	Vip            bool   `json:"vip"`
	Buildin        bool   `json:"buildin"`
	Create_user_id string `json:"create_user_id"`
	Create_time    string `json:"create_time"`
	Avatar         string `json:"avatar"`
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

func (s *service) PostUser(u []User) ([]string, error) {
	var ids []string
	for _, user := range u {
		err := s.users.Store(userToUsermodel(user))
		if err != nil {
			return ids, err
		} else {
			ids = append(ids, user.ID)
		}
	}
	return ids, nil
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

func (s *service) PutMultiUser(u []User) ([]string, error) {
	var ids []string
	for _, user := range u {
		if len(user.ID) == 0 {
			return ids, ErrInvalidArgument
		}
		_, err := s.GetUser(user.ID)
		if err != nil {
			return ids, ErrInconsistentIDs
		}
		err = s.users.Update(user.ID, userToUsermodel(user))
		if err != nil {
			return ids, err
		}
		ids = append(ids, user.ID)
	}
	return ids, nil
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

func (s *service) DeleteMultiUser(listid []string) ([]string, error) {
	var ids []string
	if len(listid) == 0 {
		return ids, ErrInvalidArgument
	}
	for _, id := range listid {
		error := s.users.Remove(id)
		if error != nil {
			return ids, ErrNotFound
		}
		ids = append(ids, id)
	}
	return ids, nil
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
