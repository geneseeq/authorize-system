// Package usering provides the use-case of booking a cargo. Used by views
// facing an administrator.
package usering

import (
	"errors"
	"time"

	"github.com/geneseeq/authorize-system/cms/user"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var (
	ErrInvalidArgument = errors.New("invalid argument")
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
	ErrExceededMount   = errors.New("exceeded max mount")
	LimitMaxSum        = 50
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
	ID           string    `json:"id"`
	Type         int       `json:"type"` //"type":"医生/教师/个人/员工/企业"
	Number       string    `json:"number"`
	Username     string    `json:"username"`
	Tele         string    `json:"tele"`
	Gneder       bool      `json:"gender"`
	Status       int       `json:"status"`
	Validity     bool      `json:"validity"`
	Vip          bool      `json:"vip"`
	Buildin      bool      `json:"buildin"`
	CreateUserID string    `json:"create_user_id"`
	CreateTime   time.Time `json:"create_time"`
	Avatar       string    `json:"avatar"`
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
	if len(u) < LimitMaxSum {
		for _, data := range u {
			data.CreateTime = user.TimeUtcToCst(time.Now())
			err := s.users.Store(userToUsermodel(data))
			if err != nil {
				return ids, err
			}
			ids = append(ids, data.ID)
		}
		return ids, nil
	}
	return ids, ErrExceededMount
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
	if len(u) < LimitMaxSum {
		for _, data := range u {
			if len(data.ID) == 0 {
				return ids, ErrInvalidArgument
			}
			_, err := s.GetUser(data.ID)
			if err != nil {
				return ids, ErrInconsistentIDs
			}
			err = s.users.Update(data.ID, userToUsermodel(data))
			if err != nil {
				return ids, err
			}
			ids = append(ids, data.ID)
		}
		return ids, nil
	}
	return ids, ErrExceededMount
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
	if len(listid) < LimitMaxSum {
		for _, id := range listid {
			error := s.users.Remove(id)
			if error != nil {
				return ids, ErrNotFound
			}
			ids = append(ids, id)
		}
		return ids, nil
	}
	return ids, ErrExceededMount
}

func userToUsermodel(u User) *user.UserModel {

	return &user.UserModel{
		ID:           u.ID,
		Type:         u.Type,
		Number:       u.Number,
		Username:     u.Username,
		Gneder:       u.Gneder,
		Tele:         u.Tele,
		Status:       u.Status,
		Validity:     u.Validity,
		Vip:          u.Vip,
		Buildin:      u.Buildin,
		CreateUserID: u.CreateUserID,
		CreateTime:   u.CreateTime,
		Avatar:       u.Avatar,
	}
}

func usermodelToUser(c *user.UserModel) User {
	return User{
		ID:           c.ID,
		Type:         c.Type,
		Number:       c.Number,
		Username:     c.Username,
		Gneder:       c.Gneder,
		Tele:         c.Tele,
		Status:       c.Status,
		Validity:     c.Validity,
		Vip:          c.Vip,
		Buildin:      c.Buildin,
		CreateUserID: c.CreateUserID,
		CreateTime:   c.CreateTime,
		Avatar:       c.Avatar,
	}
}
