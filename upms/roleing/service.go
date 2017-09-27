// Package roleing provides the use-case of booking a cargo. Used by views
// facing an administrator.
package roleing

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
	GetRole(id string) (Role, error)
	GetAllRole() ([]Role, error)
	PostRole(role []Role) ([]string, []string, error)
	DeleteRole(id string) error
	DeleteMultiRole(listid []string) ([]string, []string, error)
	PutRole(id string, role Role) error
	PutMultiRole(role []Role) ([]string, []string, error)
}

// Role is a user base info
type Role struct {
	ID           string    `json:"id"`
	Type         int       `json:"type"` //"type":"医生/教师/个人/员工/企业"
	Name         string    `json:"name"`
	Alias        string    `json:"alias"`
	Buildin      bool      `json:"buildin"`
	CreateUserID string    `json:"create_user_id"`
	CreateTime   time.Time `json:"create_time"`
	UpdateTime   time.Time `json:"update_time"`
}

type service struct {
	roles user.RoleRepository
}

// NewService creates a booking service with necessary dependencies.
func NewService(roles user.RoleRepository) Service {
	return &service{
		roles: roles,
	}
}

func (s *service) GetRole(id string) (Role, error) {
	if id == "" {
		return Role{}, ErrInvalidArgument
	}
	r, err := s.roles.Find(id)
	if err != nil {
		return Role{}, ErrNotFound
	}
	return rolemodelToRole(r), nil
}

func (s *service) GetAllRole() ([]Role, error) {
	var result []Role
	for _, g := range s.roles.FindRoleAll() {
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
			curTime := user.TimeUtcToCst(time.Now())
			role.UpdateTime = curTime
			role.CreateTime = curTime
			err := s.roles.Store(roleToRolemodel(role))
			if err != nil {
				failedIds = append(failedIds, role.ID)
				continue
			} else {
				sucessedIds = append(sucessedIds, role.ID)
			}
		}
		return sucessedIds, failedIds, nil
	}
	return nil, nil, ErrExceededMount

}

func (s *service) DeleteRole(id string) error {
	if id == "" {
		return ErrInvalidArgument
	}
	err := s.roles.Remove(id)
	if err != nil {
		return ErrNotFound
	}
	return nil
}

func (s *service) DeleteMultiRole(listid []string) ([]string, []string, error) {
	var sucessedIds []string
	var failedIds []string
	if len(listid) < LimitMaxSum {
		for _, id := range listid {
			if id == "" {
				return nil, nil, ErrInvalidArgument
			}
			err := s.roles.Remove(id)
			if err != nil {
				failedIds = append(failedIds, id)
				continue
			}
			sucessedIds = append(sucessedIds, id)
		}
		return sucessedIds, failedIds, nil
	}
	return nil, nil, ErrExceededMount
}

func (s *service) PutRole(id string, role Role) error {
	_, err := s.GetRole(id)
	if err != nil {
		return ErrInconsistentIDs
	}
	err = s.roles.Update(id, roleToRolemodel(role))
	return err
}

func (s *service) PutMultiRole(r []Role) ([]string, []string, error) {
	var sucessedIds []string
	var failedIds []string
	if len(r) < LimitMaxSum {
		for _, role := range r {
			if len(role.ID) == 0 {
				failedIds = append(failedIds, role.ID)
				continue
			}
			_, err := s.GetRole(role.ID)
			if err != nil {
				failedIds = append(failedIds, role.ID)
				continue
			}
			err = s.roles.Update(role.ID, roleToRolemodel(role))
			if err != nil {
				failedIds = append(failedIds, role.ID)
				continue
			}
			sucessedIds = append(sucessedIds, role.ID)
		}
		return sucessedIds, failedIds, nil
	}
	return nil, nil, ErrExceededMount
}

func roleToRolemodel(r Role) *user.RoleModel {

	return &user.RoleModel{
		UnionID:      r.ID,
		ID:           r.ID,
		Type:         r.Type,
		Name:         r.Name,
		Alias:        r.Alias,
		Buildin:      r.Buildin,
		CreateUserID: r.CreateUserID,
		CreateTime:   r.CreateTime,
		UpdateTime:   r.UpdateTime,
	}
}

func rolemodelToRole(r *user.RoleModel) Role {
	return Role{
		ID:           r.ID,
		Type:         r.Type,
		Name:         r.Name,
		Alias:        r.Alias,
		Buildin:      r.Buildin,
		CreateUserID: r.CreateUserID,
		CreateTime:   r.CreateTime,
		UpdateTime:   r.UpdateTime,
	}
}
