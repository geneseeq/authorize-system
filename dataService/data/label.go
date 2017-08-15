// Package data contains the heart of the domain model.
package data

import (
	"time"
)

// LabelModel is user struct
type LabelModel struct {
	LabelID    string
	SampleID   []string
	OrderID    []string
	UpdateTime time.Time
}

// NewLabelData is create instance
func NewLabelData(id string) *LabelModel {
	return &LabelModel{
		LabelID: id,
	}
}

// LableRepository is user interface
type LableRepository interface {
	Store(label *LabelModel) error
	Find(id string) (*LabelModel, error)
	FindLabelAll() []*LabelModel
	Remove(id string) error
	Update(id string, label *LabelModel) error
}
