// Package data contains the heart of the domain model.
package data

import (
	"time"
)

// FieldModel is user struct
type FieldModel struct {
	UnionID    string    `bson:"_id"` //唯一ID
	ID         string    `bson:"id"`
	Field      string    `bson:"field"`
	Type       string    `bson:"type"`
	Comment    string    `bson:"comment"`
	UpdateTime time.Time `bson:"update_time"`
	CreateTime time.Time `bson:"create_time"`
}

// NewFieldData is create instance
func NewFieldData(id string) *FieldModel {
	return &FieldModel{
		ID: id,
	}
}

// FieldRepository is user interface
type FieldRepository interface {
	Store(field *FieldModel) error
	Find(id string) (*FieldModel, error)
	FindFieldAll() []*FieldModel
	Remove(id string) error
	Update(id string, field *FieldModel) error
}
