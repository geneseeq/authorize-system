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
	PostRole(role []Role) ([]string, error)
	DeleteRole(id string) error
	// DeleteMultiGroup(listid []string) ([]string, error)
	// PutGroup(id string, group Group) error
	// PutMultiGroup(group []Group) ([]string, error)
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
	r, error := s.roles.Find(id)
	if error != nil {
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

func (s *service) PostRole(r []Role) ([]string, error) {
	var ids []string
	for _, role := range r {
		err := s.roles.Store(roleToRolemodel(role))
		if err != nil {
			return ids, err
		} else {
			ids = append(ids, role.ID)
		}
	}
	return ids, nil
}

func (s *service) DeleteRole(id string) error {
	if id == "" {
		return ErrInvalidArgument
	}
	error := s.roles.Remove(id)
	if error != nil {
		return ErrNotFound
	}
	return nil
}

// func (s *service) DeleteMultiGroup(listid []string) ([]string, error) {
// 	var ids []string
// 	if len(listid) == 0 {
// 		return ids, ErrInvalidArgument
// 	}
// 	for _, id := range listid {
// 		error := s.groups.Remove(id)
// 		if error != nil {
// 			return ids, ErrNotFound
// 		}
// 		ids = append(ids, id)
// 	}
// 	return ids, nil
// }

// func (s *service) PutGroup(id string, group Group) error {
// 	_, err := s.GetGroup(id)
// 	if err != nil {
// 		return ErrInconsistentIDs
// 	}
// 	err = s.groups.Update(id, groupToGroupmodel(group))
// 	return err
// }

// func (s *service) PutMultiGroup(g []Group) ([]string, error) {
// 	var ids []string
// 	for _, group := range g {
// 		if len(group.ID) == 0 {
// 			return ids, ErrInvalidArgument
// 		}
// 		_, err := s.GetGroup(group.ID)
// 		if err != nil {
// 			return ids, ErrInconsistentIDs
// 		}
// 		err = s.groups.Update(group.ID, groupToGroupmodel(group))
// 		if err != nil {
// 			return ids, err
// 		}
// 		ids = append(ids, group.ID)
// 	}
// 	return ids, nil
// }

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
