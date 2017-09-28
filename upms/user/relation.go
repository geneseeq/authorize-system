package user

import (
	"time"
)

// RoleID is group struct
type RoleID []string

// RoleRelationModel is group struct
type RoleRelationModel struct {
	UnionID      string    `bson:"_id"`
	UserID       string    `bson:"user_id"`
	RoleID       []string  `bson:"role_id"`
	Buildin      bool      `bson:"buildin"`
	CreateUserID string    `bson:"create_user_id"`
	CreateTime   time.Time `bson:"create_time"`
	UpdateTime   time.Time `bson:"update_time"`
}

// GroupRelationModel is group struct
type GroupRelationModel struct {
	UnionID      string    `bson:"_id"`
	GroupID      string    `bson:"group_id"`
	UserID       []string  `bson:"user_id"`
	RoleID       []string  `bson:"role_id"`
	Buildin      bool      `bson:"buildin"`
	CreateUserID string    `bson:"create_user_id"`
	CreateTime   time.Time `bson:"create_time"`
	UpdateTime   time.Time `bson:"update_time"`
}

type AuthorityModel struct {
	DataID string   `json:"data_id" bson:"data_id"`
	Action []string `json:"action" bson:"action"`
}

// AuthorityRelationModel is group struct
type AuthorityRelationModel struct {
	UnionID      string           `bson:"_id"`
	RoleID       string           `bson:"role_id"`
	Authority    []AuthorityModel `bson:"authority"`
	Validity     string           `bson:"validity"`
	Buildin      bool             `bson:"buildin"`
	CreateUserID string           `bson:"create_user_id"`
	CreateTime   time.Time        `bson:"create_time"`
	UpdateTime   time.Time        `bson:"update_time"`
}

// DeleteAuthorityModel is group struct
type DeleteAuthorityModel struct {
	RoleID string   `bson:"role_id"`
	DataID []string `bson:"data_id"`
}

// NewRoleRelation is create instance
func NewRoleRelation(id string) *RoleRelationModel {
	return &RoleRelationModel{
		UserID: id,
	}
}

// NewGroupRelation is create instance
func NewGroupRelation(id string) *GroupRelationModel {
	return &GroupRelationModel{
		GroupID: id,
	}
}

// NewAuthorityRelation is create instance
func NewAuthorityRelation(id string) *AuthorityRelationModel {
	return &AuthorityRelationModel{
		RoleID: id,
	}
}

// RelationRepository is user interface
type RelationRepository interface {
	Store(role *RoleRelationModel) error
	FindFromUser(id string) (*RoleRelationModel, error)
	FindAllFromUser() []*RoleRelationModel
	Remove(userID string, role *RoleRelationModel) error
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
