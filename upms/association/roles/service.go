// Package roles provides the use-case of booking a cargo. Used by views
// facing an administrator.
package roles

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
)

// Service is the interface that provides booking methods.
type Service interface {
	GetAuthorityFromRole(id string) (Authority, error)
	GetAllAuthority() ([]Authority, error)
	PostAuthority(role []Authority) ([]string, []string, error)
	DeleteMultiAuthority([]DeleteData) ([]string, []string, error)
	PutMultiAuthority([]Authority) ([]string, []string, error)
}

// DeleteData is ...
type DeleteData struct {
	RoleID string   `json:"role_id"`
	DataID []string `json:"data_id"`
}

// type authority struct {
// 	DataID string   `json:"data_id"`
// 	Action []string `json:"action"`
// }

// Authority is a user base authority
type Authority struct {
	ID           string                `json:"id"`
	RoleID       string                `json:"role_id"`
	Authority    []user.AuthorityModel `json:"authority"`
	Validity     string                `json:"validity"`
	Buildin      bool                  `json:"buildin"`
	CreateUserID string                `json:"create_user_id"`
	CreateTime   time.Time             `json:"create_time"`
}

type service struct {
	authoritys user.AuthorityRelationRepository
}

// NewService creates a booking service with necessary dependencies.
func NewService(authoritys user.AuthorityRelationRepository) Service {
	return &service{
		authoritys: authoritys,
	}
}

func (s *service) GetAuthorityFromRole(id string) (Authority, error) {
	if id == "" {
		return Authority{}, ErrInvalidArgument
	}
	a, err := s.authoritys.FindFromAuthority(id)
	if err != nil {
		return Authority{}, ErrNotFound
	}
	return authorityModelToAuthority(a), nil
}

func (s *service) GetAllAuthority() ([]Authority, error) {
	var result []Authority
	for _, a := range s.authoritys.FindAllFromAuthority() {
		result = append(result, authorityModelToAuthority(a))
	}
	if len(result) == 0 {
		return result, ErrNotFound
	}
	return result, nil
}

func (s *service) PostAuthority(authority []Authority) ([]string, []string, error) {
	var sucessedIds []string
	var failedIds []string
	for _, a := range authority {
		a.CreateTime = user.TimeUtcToCst(time.Now())
		err := s.authoritys.Store(authorityToAuthorityModel(a))
		if err != nil {
			failedIds = append(failedIds, a.RoleID)
			continue
		} else {
			sucessedIds = append(sucessedIds, a.RoleID)
		}
	}
	return sucessedIds, failedIds, nil
}

func (s *service) DeleteMultiAuthority(d []DeleteData) ([]string, []string, error) {
	var sucessedIds []string
	var failedIds []string
	if len(d) == 0 {
		return nil, nil, ErrInvalidArgument
	}
	for _, value := range d {
		err := s.authoritys.Remove(value.RoleID, deleteDataToDeleteModel(value))
		if err != nil {
			failedIds = append(failedIds, value.RoleID)
			continue
		}
		sucessedIds = append(sucessedIds, value.RoleID)
	}
	return sucessedIds, failedIds, nil
}

func (s *service) PutMultiAuthority(authority []Authority) ([]string, []string, error) {
	var sucessedIds []string
	var failedIds []string
	if len(authority) == 0 {
		return nil, nil, ErrInvalidArgument
	}
	for _, value := range authority {
		if len(value.RoleID) == 0 {
			failedIds = append(failedIds, value.RoleID)
			continue
		}
		err := s.authoritys.Update(value.RoleID, authorityToAuthorityModel(value))
		if err != nil {
			failedIds = append(failedIds, value.RoleID)
			continue
		}
		sucessedIds = append(sucessedIds, value.RoleID)
	}
	return sucessedIds, failedIds, nil
}

func authorityToAuthorityModel(a Authority) *user.AuthorityRelationModel {

	return &user.AuthorityRelationModel{
		ID:           a.ID,
		RoleID:       a.RoleID,
		Authority:    a.Authority,
		Validity:     a.Validity,
		Buildin:      a.Buildin,
		CreateUserID: a.CreateUserID,
		CreateTime:   a.CreateTime,
	}
}

func deleteDataToDeleteModel(d DeleteData) *user.DeleteAuthorityModel {

	return &user.DeleteAuthorityModel{
		RoleID: d.RoleID,
		DataID: d.DataID,
	}
}

func authorityModelToAuthority(a *user.AuthorityRelationModel) Authority {
	return Authority{
		ID:           a.ID,
		RoleID:       a.RoleID,
		Authority:    a.Authority,
		Validity:     a.Validity,
		Buildin:      a.Buildin,
		CreateUserID: a.CreateUserID,
		CreateTime:   a.CreateTime,
	}
}
