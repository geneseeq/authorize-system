// Package data contains the heart of the domain model.
package data

import (
	"time"
)

// FieldModel is user struct
type FieldModel struct {
	ID         string
	Field      string
	Type       string
	Comment    string
	UpdateTime time.Time
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
