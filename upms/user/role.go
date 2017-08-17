package user

import (
	"time"
)

// GroupModel is group struct
type RoleModel struct {
	ID           string
	Type         int
	Name         string
	Alias        string
	Buildin      bool
	CreateUserID string
	CreateTime   time.Time
}

// NewRole is create instance
func NewRole(id string) *RoleModel {
	return &RoleModel{
		ID: id,
	}
}

// RoleRepository is user interface
type RoleRepository interface {
	Store(role *RoleModel) error
	Find(id string) (*RoleModel, error)
	FindRoleAll() []*RoleModel
	Remove(id string) error
	Update(id string, role *RoleModel) error
}
