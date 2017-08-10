package user

import (
	"time"
)

// ServicesModel is group struct
type ServicesModel struct {
	ID           string
	Parent       string
	Depend       []string
	Name         string
	Level        string
	Path         string
	RegisterTime time.Time
	Status       string
	Validity     bool
	Buildin      bool
	Owner        []string //user type
	CreateUserID string
	CreateTime   time.Time
}

// NewService is create instance
func NewService(id string) *ServicesModel {
	return &ServicesModel{
		ID: id,
	}
}

// ServiceRepository is user interface
type ServiceRepository interface {
	Store(set *ServicesModel) error
	FindService(id string) (*ServicesModel, error)
	FindAllService() []*ServicesModel
	Remove(id string) error
	Update(id string, set *ServicesModel) error
}
