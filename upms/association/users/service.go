// Package user provides the use-case of booking a cargo. Used by views
// facing an administrator.
package users

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
	GetRoleFromUser(id string) (Role, error)
	GetAllRole() ([]Role, error)
	PostRole(role []Role) ([]string, []string, error)
	// DeleteRole(id string) error
	DeleteMultiRole([]Role) ([]string, []string, error)
	// PutRole(id string, role Role) error
	PutMultiRole([]Role) ([]string, []string, error)
}

// RoleIDList is a role id list
// type RoleIDList struct {
// 	RoleID []string
// }

// Role is a user base info
type Role struct {
	UserID       string    `json:"user_id"`
	RoleID       []string  `json:"role_id"`
	Buildin      bool      `json:"buildin"`
	CreateUserID string    `json:"create_user_id"`
	CreateTime   time.Time `json:"create_time"`
	UpdateTime   time.Time `json:"update_time"`
}

type service struct {
	roles user.RelationRepository
}

// NewService creates a booking service with necessary dependencies.
func NewService(roles user.RelationRepository) Service {
	return &service{
		roles: roles,
	}
}

func (s *service) GetRoleFromUser(id string) (Role, error) {
	if id == "" {
		return Role{}, ErrInvalidArgument
	}
	r, err := s.roles.FindFromUser(id)
	if err != nil {
		return Role{}, ErrNotFound
	}
	return rolemodelToRole(r), nil
}

func (s *service) GetAllRole() ([]Role, error) {
	var result []Role
	for _, g := range s.roles.FindAllFromUser() {
		result = append(result, rolemodelToRole(g))
	}
	if len(result) == 0 {
		return result, ErrNotFound
	}
	return result, nil
}

func (s *service) PostRole(r []Role) ([]string, []string, error) {
	var sucessedIds []string
	var failedIds []string
	if len(r) < LimitMaxSum {
		for _, role := range r {
			currentTime := user.TimeUtcToCst(time.Now())
			role.CreateTime = currentTime
			role.UpdateTime = currentTime
			err := s.roles.Store(roleToRolemodel(role))
			if err != nil {
				failedIds = append(failedIds, role.UserID)
				continue
			}
			sucessedIds = append(sucessedIds, role.UserID)
		}
		return sucessedIds, failedIds, nil
	}
	return nil, nil, ErrExceededMount
}

// func (s *service) DeleteRole(id string) error {
// 	if id == "" {
// 		return ErrInvalidArgument
// 	}
// 	err := s.roles.Remove(id)
// 	if err != nil {
// 		return ErrNotFound
// 	}
// 	return nil
// }

func (s *service) DeleteMultiRole(role []Role) ([]string, []string, error) {
	var sucessedIds []string
	var failedIds []string
	if len(role) == 0 {
		return nil, nil, ErrInvalidArgument
	}
	if len(role) < LimitMaxSum {
		for _, value := range role {
			err := s.roles.Remove(value.UserID, roleToRolemodel(value))
			if err != nil {
				failedIds = append(failedIds, value.UserID)
				continue
			}
			sucessedIds = append(sucessedIds, value.UserID)
		}
		return sucessedIds, failedIds, nil
	}
	return nil, nil, ErrExceededMount
}

// func (s *service) PutRole(id string, role Role) error {
// 	_, err := s.GetRole(id)
// 	if err != nil {
// 		return ErrInconsistentIDs
// 	}
// 	err = s.roles.Update(id, roleToRolemodel(role))
// 	return err
// }

func (s *service) PutMultiRole(role []Role) ([]string, []string, error) {
	var sucessedIds []string
	var failedIds []string
	if len(role) == 0 {
		return nil, nil, ErrInvalidArgument
	}
	if len(role) < LimitMaxSum {
		for _, value := range role {
			if len(value.UserID) == 0 {
				failedIds = append(failedIds, value.UserID)
				continue
			}
			err := s.roles.Update(value.UserID, roleToRolemodel(value))
			if err != nil {
				failedIds = append(failedIds, value.UserID)
				continue
			}
			sucessedIds = append(sucessedIds, value.UserID)
		}
		return sucessedIds, failedIds, nil
	}
	return nil, nil, ErrExceededMount
}

func roleToRolemodel(r Role) *user.RoleRelationModel {

	return &user.RoleRelationModel{
		UnionID:      r.UserID,
		UserID:       r.UserID,
		RoleID:       r.RoleID,
		Buildin:      r.Buildin,
		CreateUserID: r.CreateUserID,
		CreateTime:   r.CreateTime,
		UpdateTime:   r.UpdateTime,
	}
}

func rolemodelToRole(r *user.RoleRelationModel) Role {
	return Role{
		UserID:       r.UserID,
		RoleID:       r.RoleID,
		Buildin:      r.Buildin,
		CreateUserID: r.CreateUserID,
		CreateTime:   r.CreateTime,
		UpdateTime:   r.UpdateTime,
	}
}
