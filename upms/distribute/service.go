// Package distribute provides the use-case of booking a cargo. Used by views
// facing an administrator.
package distribute

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
	GetRoleDistribute(id string) (RoleDistribute, error)
	GetAllRoleDistribute() ([]RoleDistribute, error)
	PostRoleDistribute(distribute []RoleDistribute) ([]string, []string, error)
	DeleteMultiRoleDistribute(listid []string) ([]string, []string, error)
	PutMultiRoleDistribute(distribute []RoleDistribute) ([]string, []string, error)
}

// Distribute is a user base info
type RoleDistribute struct {
	ID           string    `json:"id"`
	GroupID      string    `json:"group_id"` //组类型（所有具备父子关系的组：机构，部门，项目组）
	UserID       string    `json:"user_id"`  //组织编码
	RoleID       string    `json:"role_id"`
	Buildin      bool      `json:"buildin"`        //是否内建，true内建，false非内建
	CreateUserID string    `json:"create_user_id"` //创建人ID
	UpdateUserID string    `json:"update_user_id"`
	CreateTime   time.Time `json:"create_time"` //创建时间
	UpdateTime   time.Time `json:"update_time"` //更新时间
}

type service struct {
	distribute user.RoleDistributeRepository
}

// NewService creates a booking service with necessary dependencies.
func NewService(distribute user.RoleDistributeRepository) Service {
	return &service{
		distribute: distribute,
	}
}

func (s *service) GetRoleDistribute(id string) (RoleDistribute, error) {
	if id == "" {
		return RoleDistribute{}, ErrInvalidArgument
	}
	c, error := s.distribute.Find(id)
	if error != nil {
		return RoleDistribute{}, ErrNotFound
	}
	return roleDistributeModelToRoleDistribute(c), nil
}

func (s *service) GetAllRoleDistribute() ([]RoleDistribute, error) {
	var result []RoleDistribute
	for _, g := range s.distribute.FindRoleDistributeAll() {
		result = append(result, roleDistributeModelToRoleDistribute(g))
	}
	return result, nil
}

func (s *service) PostRoleDistribute(r []RoleDistribute) ([]string, []string, error) {
	var sucessed []string
	var failed []string
	if len(r) < LimitMaxSum {
		for _, distribute := range r {
			curTime := user.TimeUtcToCst(time.Now())
			distribute.CreateTime = curTime
			distribute.UpdateTime = curTime
			err := s.distribute.Store(roleDistributeToRoleDistributeModel(distribute))
			if err != nil {
				failed = append(failed, distribute.ID)
				continue
			}
			sucessed = append(sucessed, distribute.ID)
		}
		return sucessed, failed, nil
	}
	return nil, nil, ErrExceededMount
}

func (s *service) DeleteMultiRoleDistribute(listid []string) ([]string, []string, error) {
	var sucessed []string
	var failed []string
	if len(listid) < LimitMaxSum {
		for _, id := range listid {
			if id == "" {
				return nil, nil, ErrInvalidArgument
			}
			error := s.distribute.Remove(id)
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

func (s *service) PutMultiRoleDistribute(r []RoleDistribute) ([]string, []string, error) {
	var sucessed []string
	var failed []string
	if len(r) < LimitMaxSum {
		for _, distribute := range r {
			if len(distribute.ID) == 0 {
				failed = append(failed, distribute.ID)
				continue
			}
			_, err := s.GetRoleDistribute(distribute.ID)
			if err != nil {
				failed = append(failed, distribute.ID)
				continue
			}
			err = s.distribute.Update(distribute.ID, roleDistributeToRoleDistributeModel(distribute))
			if err != nil {
				failed = append(failed, distribute.ID)
				continue
			}
			sucessed = append(sucessed, distribute.ID)
		}
		return sucessed, failed, nil
	}
	return nil, nil, ErrExceededMount
}

func roleDistributeToRoleDistributeModel(g RoleDistribute) *user.RoleDistributeModel {
	return &user.RoleDistributeModel{
		UnionID:      g.ID,
		ID:           g.ID,
		GroupID:      g.GroupID,
		UserID:       g.UserID,
		RoleID:       g.RoleID,
		Buildin:      g.Buildin,
		CreateUserID: g.CreateUserID,
		UpdateUserID: g.UpdateUserID,
		CreateTime:   g.CreateTime,
		UpdateTime:   g.UpdateTime,
	}
}

func roleDistributeModelToRoleDistribute(g *user.RoleDistributeModel) RoleDistribute {
	return RoleDistribute{
		ID:           g.ID,
		GroupID:      g.GroupID,
		UserID:       g.UserID,
		RoleID:       g.RoleID,
		Buildin:      g.Buildin,
		CreateUserID: g.CreateUserID,
		UpdateUserID: g.UpdateUserID,
		CreateTime:   g.CreateTime,
		UpdateTime:   g.UpdateTime,
	}
}
