// Package roleing provides the use-case of booking a cargo. Used by views
// facing an administrator.
package roleing

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
	ID             string `json:"id"`
	Type           int    `json:"type"` //"type":"医生/教师/个人/员工/企业"
	Parent         string `json:"parent"`
	Name           string `json:"name"`
	Code           string `json:"code"`
	Alias          string `json:"alias"`
	Buildin        bool   `json:"buildin"`
	Create_user_id string `json:"create_user_id"`
	Create_time    string `json:"create_time"`
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
	for _, role := range r {
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
	if len(listid) == 0 {
		return nil, nil, ErrInvalidArgument
	}
	for _, id := range listid {
		err := s.roles.Remove(id)
		if err != nil {
			failedIds = append(failedIds, id)
			continue
		}
		sucessedIds = append(sucessedIds, id)
	}
	return sucessedIds, failedIds, nil
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
			continue
		}
		sucessedIds = append(sucessedIds, role.ID)
	}
	return sucessedIds, failedIds, nil
}

func roleToRolemodel(g Role) *user.RoleModel {

	return &user.RoleModel{
		ID:             g.ID,
		Type:           g.Type,
		Parent:         g.Parent,
		Name:           g.Name,
		Code:           g.Code,
		Alias:          g.Alias,
		Buildin:        g.Buildin,
		Create_user_id: g.Create_user_id,
		Create_time:    g.Create_time,
	}
}

func rolemodelToRole(g *user.RoleModel) Role {
	return Role{
		ID:             g.ID,
		Type:           g.Type,
		Parent:         g.Parent,
		Name:           g.Name,
		Code:           g.Code,
		Alias:          g.Alias,
		Buildin:        g.Buildin,
		Create_user_id: g.Create_user_id,
		Create_time:    g.Create_time,
	}
}
