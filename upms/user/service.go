package user

import (
	"time"
)

// ServicesModel is group struct
type ServicesModel struct {
	UnionID      string    `bson:"_id"` //唯一ID
	ID           string    `bson:"id"`
	Parent       string    `bson:"parent"`
	Depend       []string  `bson:"depend"`
	Name         string    `bson:"name"`
	Level        string    `bson:"level"`
	Path         string    `bson:"path"`
	RegisterTime time.Time `bson:"register_time"`
	Status       string    `bson:"status"`
	Validity     bool      `bson:"validity"`
	Buildin      bool      `bson:"buildin"`
	Owner        []string  `bson:"owner"` //user type
	UpdateUserID string    `bson:"update_user_id"`
	CreateUserID string    `bson:"create_user_id"`
	CreateTime   time.Time `bson:"create_time"`
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
