package user

import "time"

// GroupModel is group struct
type GroupModel struct {
	Parent       string
	ID           string
	Type         int
	Name         string
	Code         string
	Alias        string
	Buildin      bool
	CreateUserID string
	CreateTime   time.Time
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