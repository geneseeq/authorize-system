// Package grouping provides the use-case of booking a cargo. Used by views
// facing an administrator.
package grouping

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/geneseeq/authorize-system/upms/user"
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
	GetGroup(id string) (Group, error)
	GetAllGroup() ([]Group, error)
	PostGroup(group []Group) ([]string, error)
	DeleteGroup(id string) error
	DeleteMultiGroup(listid []string) ([]string, error)
	PutGroup(id string, group Group) error
	PutMultiGroup(group []Group) ([]string, error)
}

// Group is a user base info
type Group struct {
	ID           string    `json:"id"`
	Type         int       `json:"type"` //"type":"医生/教师/个人/员工/企业"
	Parent       string    `json:"parent"`
	Name         string    `json:"name"`
	Code         string    `json:"code"`
	Alias        string    `json:"alias"`
	Buildin      bool      `json:"buildin"`
	CreateUserID string    `json:"create_user_id"`
	CreateTime   time.Time `json:"create_time"`
}

type service struct {
	groups user.GroupRepository
}

// NewService creates a booking service with necessary dependencies.
func NewService(groups user.GroupRepository) Service {
	return &service{
		groups: groups,
	}
}

func (s *service) GetGroup(id string) (Group, error) {
	if id == "" {
		return Group{}, ErrInvalidArgument
	}
	c, error := s.groups.Find(id)
	if error != nil {
		return Group{}, ErrNotFound
	}
	return groupmodelToGroup(c), nil
}

func (s *service) GetAllGroup() ([]Group, error) {
	var result []Group
	for _, g := range s.groups.FindGroupAll() {
		result = append(result, groupmodelToGroup(g))
	}
	return result, nil
}

func (s *service) PostGroup(g []Group) ([]string, error) {
	var ids []string
	for _, group := range g {
		group.CreateTime = user.TimeUtcToCst(time.Now())
		err := s.groups.Store(groupToGroupmodel(group))
		if err != nil {
			return ids, err
		} else {
			ids = append(ids, group.ID)
		}
	}
	return ids, nil
}

func (s *service) DeleteGroup(id string) error {
	if id == "" {
		return ErrInvalidArgument
	}
	error := s.groups.Remove(id)
	if error != nil {
		return ErrNotFound
	}
	return nil
}

func (s *service) DeleteMultiGroup(listid []string) ([]string, error) {
	var ids []string
	if len(listid) == 0 {
		return ids, ErrInvalidArgument
	}
	for _, id := range listid {
		error := s.groups.Remove(id)
		if error != nil {
			return ids, ErrNotFound
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (s *service) PutGroup(id string, group Group) error {
	_, err := s.GetGroup(id)
	if err != nil {
		return ErrInconsistentIDs
	}
	err = s.groups.Update(id, groupToGroupmodel(group))
	return err
}

func (s *service) PutMultiGroup(g []Group) ([]string, error) {
	var ids []string
	for _, group := range g {
		if len(group.ID) == 0 {
			return ids, ErrInvalidArgument
		}
		_, err := s.GetGroup(group.ID)
		if err != nil {
			return ids, ErrInconsistentIDs
		}
		err = s.groups.Update(group.ID, groupToGroupmodel(group))
		if err != nil {
			return ids, err
		}
		ids = append(ids, group.ID)
	}
	return ids, nil
}

func groupToGroupmodel(g Group) *user.GroupModel {

	return &user.GroupModel{
		UnionID:      bson.ObjectIdHex(g.ID),
		ID:           g.ID,
		Type:         g.Type,
		Parent:       g.Parent,
		Name:         g.Name,
		Code:         g.Code,
		Alias:        g.Alias,
		Buildin:      g.Buildin,
		CreateUserID: g.CreateUserID,
		CreateTime:   g.CreateTime,
	}
}

func groupmodelToGroup(g *user.GroupModel) Group {
	return Group{
		ID:           g.ID,
		Type:         g.Type,
		Parent:       g.Parent,
		Name:         g.Name,
		Code:         g.Code,
		Alias:        g.Alias,
		Buildin:      g.Buildin,
		CreateUserID: g.CreateUserID,
		CreateTime:   g.CreateTime,
	}
}
