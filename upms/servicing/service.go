// Package servicing provides the use-case of booking a cargo. Used by views
// facing an administrator.
package servicing

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
	GetService(id string) (Services, error)
	PostService(service []Services) ([]string, []string, error)
	GetAllService() ([]Services, error)
	DeleteMultiService(listid []string) ([]string, []string, error)
	PutMultiService(service []Services) ([]string, []string, error)
}

// Services is a service base info
type Services struct {
	ID           string    `json:"id"`
	Parent       string    `json:"parent"` //"type":"医生/教师/个人/员工/企业"
	Depend       []string  `json:"depend"`
	Name         string    `json:"name"`
	Level        string    `json:"level"`
	Path         string    `json:"path"`
	RegisterTime time.Time `json:"register_time"`
	Status       string    `json:"status"`
	Validity     bool      `json:"validity"`
	Buildin      bool      `json:"buildin"`
	CreateUserID string    `json:"create_user_id"`
	CreateTime   time.Time `json:"create_time"`
	Owner        []string  `json:"owner"`
}

type service struct {
	services user.ServiceRepository
}

// NewService creates a booking service with necessary dependencies.
func NewService(services user.ServiceRepository) Service {
	return &service{
		services: services,
	}
}

func (s *service) PostService(u []Services) ([]string, []string, error) {
	var sucessed []string
	var failed []string
	if len(u) < LimitMaxSum {
		for _, data := range u {
			data.CreateTime = user.TimeUtcToCst(time.Now())
			err := s.services.Store(serviceToServiceModel(data))
			if err != nil {
				failed = append(failed, data.ID)
				continue
			}
			sucessed = append(sucessed, data.ID)
		}
		return sucessed, failed, nil
	}
	return sucessed, failed, ErrExceededMount
}

func (s *service) GetService(id string) (Services, error) {
	if id == "" {
		return Services{}, ErrInvalidArgument
	}
	c, error := s.services.FindService(id)
	if error != nil {
		return Services{}, ErrNotFound
	}
	return serviceModelToService(c), nil
}

func (s *service) GetAllService() ([]Services, error) {
	var result []Services
	for _, c := range s.services.FindAllService() {
		result = append(result, serviceModelToService(c))
	}
	return result, nil
}

func (s *service) PutMultiService(u []Services) ([]string, []string, error) {
	var sucessed []string
	var failed []string
	if len(u) < LimitMaxSum {
		for _, data := range u {
			if len(data.ID) == 0 {
				failed = append(failed, data.ID)
				continue
			}
			_, err := s.GetService(data.ID)
			if err != nil {
				failed = append(failed, data.ID)
				continue
			}
			err = s.services.Update(data.ID, serviceToServiceModel(data))
			if err != nil {
				failed = append(failed, data.ID)
				continue
			}
			sucessed = append(sucessed, data.ID)
		}
		return sucessed, failed, nil
	}
	return sucessed, failed, ErrExceededMount
}

func (s *service) DeleteMultiService(listid []string) ([]string, []string, error) {
	var sucessed []string
	var failed []string
	if len(listid) < LimitMaxSum {
		for _, id := range listid {
			error := s.services.Remove(id)
			if error != nil {
				failed = append(failed, id)
				continue
			}
			sucessed = append(sucessed, id)
		}
		return sucessed, failed, nil
	}
	return sucessed, failed, ErrExceededMount
}

func serviceToServiceModel(s Services) *user.ServicesModel {

	return &user.ServicesModel{
		ID:           s.ID,
		Parent:       s.Parent,
		Buildin:      s.Buildin,
		CreateTime:   s.CreateTime,
		CreateUserID: s.CreateUserID,
		Depend:       s.Depend,
		Level:        s.Level,
		Name:         s.Name,
		Owner:        s.Owner,
		Path:         s.Path,
		RegisterTime: s.RegisterTime,
		Status:       s.Status,
		Validity:     s.Validity,
	}
}

func serviceModelToService(s *user.ServicesModel) Services {
	return Services{
		ID:           s.ID,
		Parent:       s.Parent,
		Buildin:      s.Buildin,
		CreateTime:   s.CreateTime,
		CreateUserID: s.CreateUserID,
		Depend:       s.Depend,
		Level:        s.Level,
		Name:         s.Name,
		Owner:        s.Owner,
		Path:         s.Path,
		RegisterTime: s.RegisterTime,
		Status:       s.Status,
		Validity:     s.Validity,
	}
}
