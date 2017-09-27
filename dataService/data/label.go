// Package data contains the heart of the domain model.
package data

import (
	"time"
)

// LabelModel is user struct
type LabelModel struct {
	UnionID    string    `bson:"_id"` //唯一ID
	LabelID    string    `bson:"label_id"`
	SampleID   []string  `bson:"sample_id"`
	OrderID    []string  `bson:"order_id"`
	MedicalID  []string  `bson:"medical_id"`
	CreateTime time.Time `bson:"create_time"`
	UpdateTime time.Time `bson:"update_time"`
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
