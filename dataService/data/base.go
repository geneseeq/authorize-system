// Package data contains the heart of the domain model.
package data

import (
	"errors"
	"time"
)

// LabelIDModel is
type LabelIDModel struct {
	ID      string   `json:"id"`
	LabelID []string `json:"label_id"`
}

// BaseDataModel is user struct
type BaseDataModel struct {
	ID           string
	SampleID     string
	OrderID      string
	SaleID       string
	Doctor       string
	Hospital     string
	HospitalDept string
	School       string
	SchoolDept   string
	Product      string
	Project      string
	LabelID      []string
	UpdateTime   time.Time
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
