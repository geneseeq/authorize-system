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

// GroupRelationModel is group struct
type GroupRelationModel struct {
	ID           string
	GroupID      string
	UserID       []string
	RoleID       []string
	Buildin      bool
	CreateUserID string
	CreateTime   time.Time
}

type AuthorityModel struct {
	DataID string   `json:"data_id"`
	Action []string `json:"action"`
}

// AuthorityRelationModel is group struct
type AuthorityRelationModel struct {
	ID           string
	RoleID       string
	Authority    []AuthorityModel
	Validity     string
	Buildin      bool
	CreateUserID string
	CreateTime   time.Time
}

// DeleteAuthorityModel is group struct
type DeleteAuthorityModel struct {
	RoleID string
	DataID []string
}

// NewRoleRelation is create instance
func NewRoleRelation(id string) *RoleRelationModel {
	return &RoleRelationModel{
		ID: id,
	}
}

// NewGroupRelation is create instance
func NewGroupRelation(id string) *GroupRelationModel {
	return &GroupRelationModel{
		ID: id,
	}
}

// NewAuthorityRelation is create instance
func NewAuthorityRelation(id string) *AuthorityRelationModel {
	return &AuthorityRelationModel{
		ID: id,
	}
}

// RelationRepository is user interface
type RelationRepository interface {
	Store(role *RoleRelationModel) error
	FindFromUser(id string) (*RoleRelationModel, error)
	FindAllFromUser() []*RoleRelationModel
	Remove(userID string, roleID []string) error
	Update(id string, role *RoleRelationModel) error
}

// GroupRelationRepository is user interface
type GroupRelationRepository interface {
	Store(group *GroupRelationModel) error
	FindFromGroup(id string) (*GroupRelationModel, error)
	FindAllFromGroup() []*GroupRelationModel
	Remove(id string, group *GroupRelationModel) error
	Update(id string, group *GroupRelationModel) error
}

// AuthorityRelationRepository is user interface
type AuthorityRelationRepository interface {
	Store(authority *AuthorityRelationModel) error
	FindFromAuthority(id string) (*AuthorityRelationModel, error)
	FindAllFromAuthority() []*AuthorityRelationModel
	Remove(id string, d *DeleteAuthorityModel) error
	Update(id string, authority *AuthorityRelationModel) error
}
