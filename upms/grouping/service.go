// Package grouping provides the use-case of booking a cargo. Used by views
// facing an administrator.
package grouping

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
	GetGroup(id string) (Group, error)
	GetAllGroup() ([]Group, error)
	PostGroup(group []Group) ([]string, []string, error)
	DeleteGroup(id string) error
	DeleteMultiGroup(listid []string) ([]string, []string, error)
	PutGroup(id string, group Group) error
	PutMultiGroup(group []Group) ([]string, []string, error)
}

// Group is a user base info
type Group struct {
	ID           string    `json:"id"`             //用户组ID
	Type         int       `json:"type"`           //"1":"医生","2":"教师","3":"个人","4":"内部员工","0":"外部企业"
	Parent       string    `json:"parent"`         //用户父组ID
	Name         string    `json:"name"`           //组名字
	Code         string    `json:"code"`           //组织编码
	Alias        string    `json:"alias"`          //组别名
	Buildin      bool      `json:"buildin"`        //是否内建，true内建，false非内建
	CreateUserID string    `json:"create_user_id"` //创建人ID
	CreateTime   time.Time `json:"create_time"`    //创建时间
	UpdateTime   time.Time `json:"update_time"`    //更新时间
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

func (s *service) PostGroup(g []Group) ([]string, []string, error) {
	var sucessed []string
	var failed []string
	if len(g) < LimitMaxSum {
		for _, group := range g {
			curTime := user.TimeUtcToCst(time.Now())
			group.CreateTime = curTime
			group.UpdateTime = curTime
			err := s.groups.Store(groupToGroupmodel(group))
			if err != nil {
				failed = append(failed, group.ID)
				continue
			}
			sucessed = append(sucessed, group.ID)
		}
		return sucessed, failed, nil
	}
	return nil, nil, ErrExceededMount
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

func (s *service) DeleteMultiGroup(listid []string) ([]string, []string, error) {
	var sucessed []string
	var failed []string
	if len(listid) < LimitMaxSum {
		for _, id := range listid {
			if id == "" {
				return nil, nil, ErrInvalidArgument
			}
			error := s.groups.Remove(id)
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

func (s *service) PutGroup(id string, group Group) error {
	_, err := s.GetGroup(id)
	if err != nil {
		return ErrInconsistentIDs
	}
	err = s.groups.Update(id, groupToGroupmodel(group))
	return err
}

func (s *service) PutMultiGroup(g []Group) ([]string, []string, error) {
	var sucessed []string
	var failed []string
	if len(g) < LimitMaxSum {
		for _, group := range g {
			if len(group.ID) == 0 {
				failed = append(failed, group.ID)
				continue
			}
			_, err := s.GetGroup(group.ID)
			if err != nil {
				failed = append(failed, group.ID)
				continue
			}
			err = s.groups.Update(group.ID, groupToGroupmodel(group))
			if err != nil {
				failed = append(failed, group.ID)
				continue
			}
			sucessed = append(sucessed, group.ID)
		}
		return sucessed, failed, nil
	}
	return nil, nil, ErrExceededMount
}

func groupToGroupmodel(g Group) *user.GroupModel {
	return &user.GroupModel{
		UnionID:      g.ID,
		ID:           g.ID,
		Type:         g.Type,
		Parent:       g.Parent,
		Name:         g.Name,
		Code:         g.Code,
		Alias:        g.Alias,
		Buildin:      g.Buildin,
		CreateUserID: g.CreateUserID,
		CreateTime:   g.CreateTime,
		UpdateTime:   g.UpdateTime,
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
		UpdateTime:   g.UpdateTime,
	}
}
