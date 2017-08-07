package user

import (
	"time"
)

// RoleID is group struct
type RoleID []string

// RoleRelationModel is group struct
type RoleRelationModel struct {
	ID           string
	UserID       string
	RoleID       []string
	Buildin      bool
	CreateUserID string
	CreateTime   time.Time
}

// NewRoleRelation is create instance
func NewRoleRelation(id string) *RoleRelationModel {
	return &RoleRelationModel{
		ID: id,
	}
}

// RelationRepository is user interface
type RelationRepository interface {
	Store(role *RoleRelationModel) error
	FindFromUser(id string) (*RoleRelationModel, error)
	FindAllFromUser() []*RoleRelationModel
	Remove(user_id string, role_id []string) error
	Update(id string, role *RoleRelationModel) error
}
