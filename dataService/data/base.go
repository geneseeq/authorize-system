// Package data contains the heart of the domain model.
package data

import (
	"errors"
	"time"
)

// LabelIDModel is
type LabelIDModel struct {
	ID      string   `json:"id" bson:"id"`
	LabelID []string `json:"label_id" bson:"label_id"`
}

// BaseDataModel is user struct
type BaseDataModel struct {
	UnionID      string    `bson:"_id"` //唯一ID
	ID           string    `bson:"id"`
	SampleID     string    `bson:"sample_id"`
	OrderID      string    `bson:"order_id"`
	SaleID       string    `bson:"sale_id"`
	Doctor       string    `bson:"doctor"`
	Hospital     string    `bson:"hospital"`
	HospitalDept string    `bson:"hospital_dept"`
	School       string    `bson:"school"`
	SchoolDept   string    `bson:"school_dept"`
	Product      string    `bson:"product"`
	Project      string    `bson:"project"`
	LabelID      []string  `bson:"label_id"`
	UpdateTime   time.Time `bson:"update_time"`
	CreateTime   time.Time `bson:"create_time"`
}

// NewBaseData is create instance
func NewBaseData(id string) *BaseDataModel {
	return &BaseDataModel{
		ID: id,
	}
}

// ErrUnknown is unkown user error
var (
	ErrUnknown = errors.New("unknown user")
)

// TimeUtcToCst is format time
func TimeUtcToCst(t time.Time) time.Time {
	return t.Add(time.Hour * time.Duration(8))
}

// Repository is user interface
type DataRepository interface {
	Store(data *BaseDataModel) error
	Find(id string) (*BaseDataModel, error)
	FindDataAll() []*BaseDataModel
	Remove(id string) error
	RemoveLabel(id string, label []string) error
	Update(id string, data *BaseDataModel) error
}
