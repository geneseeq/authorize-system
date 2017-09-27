package user

import (
	"time"
)

// GroupModel is group struct
type RoleModel struct {
	UnionID      string    `bson:"_id"` //唯一ID
	ID           string    `bson:"id"`
	Type         int       `bson:"type"`
	Name         string    `bson:"name"`
	Alias        string    `bson:"alias"`
	Buildin      bool      `bson:"buildin"`
	CreateUserID string    `bson:"create_user_id"`
	CreateTime   time.Time `bson:"create_time"`
	UpdateTime   time.Time `bson:"update_time"`
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
