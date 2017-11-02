package user

import (
	"time"
)

// GroupModel is group struct
type GroupModel struct {
	UnionID      string    `bson:"_id"` //唯一ID
	Parent       string    `bson:"parent"`
	ID           string    `bson:"id"`
	Type         string    `bson:"type"`
	Name         string    `bson:"name"`
	Code         string    `bson:"code"`
	Alias        string    `bson:"alias"`
	Buildin      bool      `bson:"buildin"`
	CreateUserID string    `bson:"create_user_id"`
	UpdateUserID string    `bson:"update_user_id"`
	Validity     bool      `bson:"validity"`
	CreateTime   time.Time `bson:"create_time"`
	UpdateTime   time.Time `bson:"update_time"`
}

// NewGroup is create instance
func NewGroup(id string) *GroupModel {
	return &GroupModel{
		ID: id,
	}
}

// GroupRepository is user interface
type GroupRepository interface {
	Store(group *GroupModel) error
	Find(id string) (*GroupModel, error)
	FindGroupAll() []*GroupModel
	Remove(id string) error
	Update(id string, group *GroupModel) error
}
