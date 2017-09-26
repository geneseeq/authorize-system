// Package usering provides the use-case of booking a cargo. Used by views
// facing an administrator.
package usering

import (
	"errors"
	"time"

	"github.com/geneseeq/authorize-system/upms/user"
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
	PostUser(user []User) ([]string, []string, error)
	GetAllUser() ([]User, error)
	PutUser(id string, user User) error
	DeleteUser(id string) error
	DeleteMultiUser(listid []string) ([]string, []string, error)
	PutMultiUser(user []User) ([]string, []string, error)
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
	UpdateTime   time.Time `json:"update_time"`
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

func (s *service) PostUser(u []User) ([]string, []string, error) {
	var sucessed []string
	var failed []string
	if len(u) < LimitMaxSum {
		for _, data := range u {
			curTime := user.TimeUtcToCst(time.Now())
			data.CreateTime = curTime
			data.UpdateTime = curTime
			err := s.users.Store(userToUsermodel(data))
			if err != nil {
				failed = append(failed, data.ID)
				continue
			}
			sucessed = append(sucessed, data.ID)
		}
		return sucessed, failed, nil
	}
	return nil, nil, ErrExceededMount
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

func (s *service) PutUser(id string, u User) error {
	_, err := s.GetUser(id)
	if err != nil {
		return ErrInconsistentIDs
	}
	err = s.users.Update(id, userToUsermodel(u))
	return err
}

func (s *service) PutMultiUser(u []User) ([]string, []string, error) {
	var sucessed []string
	var failed []string
	if len(u) < LimitMaxSum {
		for _, data := range u {
			if len(data.ID) == 0 {
				return nil, nil, ErrInvalidArgument
			}
			_, err := s.GetUser(data.ID)
			if err != nil {
				failed = append(failed, data.ID)
				continue
			}
			err = s.users.Update(data.ID, userToUsermodel(data))
			if err != nil {
				failed = append(failed, data.ID)
				continue
			}
			sucessed = append(sucessed, data.ID)
		}
		return sucessed, failed, nil
	}
	return nil, nil, ErrExceededMount
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

func (s *service) DeleteMultiUser(listid []string) ([]string, []string, error) {
	var sucessed []string
	var failed []string
	if len(listid) < LimitMaxSum {
		for _, id := range listid {
			if id == "" {
				return nil, nil, ErrInvalidArgument
			}
			error := s.users.Remove(id)
			if error != nil {
				failed = append(failed, id)
				continue
			}
			sucessed = append(sucessed, id)
		}
		return sucessed, failed, nil
	}
	return nil, nil, ErrExceededMount
}

func userToUsermodel(u User) *user.UserModel {

	return &user.UserModel{
		UnionID:      u.ID,
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
		UpdateTime:   u.UpdateTime,
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
		UpdateTime:   c.UpdateTime,
		Avatar:       c.Avatar,
	}
}
