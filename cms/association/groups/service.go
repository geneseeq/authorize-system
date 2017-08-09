// Package groups provides the use-case of booking a cargo. Used by views
package groups

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
)

// Service is the interface that provides booking methods.
type Service interface {
	GetDataFromGroup(id string) (Groups, error)
	GetAllData() ([]Groups, error)
	PostData(role []Groups) ([]string, []string, error)
	DeleteMultiData([]Groups) ([]string, []string, error)
	PutMultiData([]Groups) ([]string, []string, error)
}

// Groups is a relation coll
type Groups struct {
	ID           string    `json:"id"`
	GroupID      string    `json:"group_id"`
	RoleID       []string  `json:"role_id"`
	UserID       []string  `json:"user_id"`
	Buildin      bool      `json:"buildin"`
	CreateUserID string    `json:"create_user_id"`
	CreateTime   time.Time `json:"create_time"`
}

type service struct {
	groups user.GroupRelationRepository
}

// NewService creates a booking service with necessary dependencies.
func NewService(groups user.GroupRelationRepository) Service {
	return &service{
		groups: groups,
	}
}

func (s *service) GetDataFromGroup(id string) (Groups, error) {
	if id == "" {
		return Groups{}, ErrInvalidArgument
	}
	g, err := s.groups.FindFromGroup(id)
	if err != nil {
		return Groups{}, ErrNotFound
	}
	return groupmodelToGroup(g), nil
}

func (s *service) GetAllData() ([]Groups, error) {
	var result []Groups
	for _, g := range s.groups.FindAllFromGroup() {
		result = append(result, groupmodelToGroup(g))
	}
	if len(result) == 0 {
		return result, ErrNotFound
	}
	return result, nil
}

func (s *service) PostData(g []Groups) ([]string, []string, error) {
	var sucessedIds []string
	var failedIds []string
	for _, group := range g {
		group.CreateTime = user.TimeUtcToCst(time.Now())
		err := s.groups.Store(groupToGroupmodel(group))
		if err != nil {
			failedIds = append(failedIds, group.GroupID)
			continue
		} else {
			sucessedIds = append(sucessedIds, group.GroupID)
		}
	}
	return sucessedIds, failedIds, nil
}

func (s *service) DeleteMultiData(g []Groups) ([]string, []string, error) {
	var sucessedIds []string
	var failedIds []string
	if len(g) == 0 {
		return nil, nil, ErrInvalidArgument
	}
	for _, value := range g {
		err := s.groups.Remove(value.GroupID, groupToGroupmodel(value))
		if err != nil {
			failedIds = append(failedIds, value.GroupID)
			continue
		}
		sucessedIds = append(sucessedIds, value.GroupID)
	}
	return sucessedIds, failedIds, nil
}

func (s *service) PutMultiData(g []Groups) ([]string, []string, error) {
	var sucessedIds []string
	var failedIds []string
	if len(g) == 0 {
		return nil, nil, ErrInvalidArgument
	}
	for _, value := range g {
		if len(value.GroupID) == 0 {
			failedIds = append(failedIds, value.GroupID)
			continue
		}
		err := s.groups.Update(value.GroupID, groupToGroupmodel(value))
		if err != nil {
			failedIds = append(failedIds, value.GroupID)
			continue
		}
		sucessedIds = append(sucessedIds, value.GroupID)
	}
	return sucessedIds, failedIds, nil
}

func groupToGroupmodel(g Groups) *user.GroupRelationModel {

	return &user.GroupRelationModel{
		ID:           g.ID,
		GroupID:      g.GroupID,
		UserID:       g.UserID,
		RoleID:       g.RoleID,
		Buildin:      g.Buildin,
		CreateUserID: g.CreateUserID,
		CreateTime:   g.CreateTime,
	}
}

func groupmodelToGroup(g *user.GroupRelationModel) Groups {
	return Groups{
		ID:           g.ID,
		GroupID:      g.GroupID,
		UserID:       g.UserID,
		RoleID:       g.RoleID,
		Buildin:      g.Buildin,
		CreateUserID: g.CreateUserID,
		CreateTime:   g.CreateTime,
	}
}
