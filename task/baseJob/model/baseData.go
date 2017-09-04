// Package model contains the heart of the domain model.
package model

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// BaseInfoModel is user struct
type BaseInfoModel struct {
	ID           bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	SampleID     string        `json:"sample_id" bson:"SAMPLE_CODE"`
	LabID        string        `json:"lab_id" bson:"LAB_CODE"`
	OrderID      string        `json:"order_id" bson:"ORDER_CODE"`
	SaleID       string        `json:"sale_id" bson:"COMMISSIONER"`
	Doctor       string        `json:"doctor" bson:"ATTENDING_DOCTOR"`
	Hospital     string        `json:"hospital" bson:"MEDICAL_INSTITUTIONS"`
	HospitalDept string        `json:"hospital_dept"`
	School       string        `json:"school"`
	SchoolDept   string        `json:"school_dept"`
	Product      string        `json:"product" bson:"PRODUCT_ID"`
	Project      string        `json:"project"`
	LabelID      []string      `json:"label_id"`
	UpdateTime   time.Time     `json:"update_time"`
}

// NewBaseInfoModel is create instance
func NewBaseInfoModel(id bson.ObjectId) *BaseInfoModel {
	return &BaseInfoModel{
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

// BaseRepository is user interface
type BaseRepository interface {
	Distinct(string, bson.M) ([]string, error)
	Aggregate(*[]bson.M) ([]BaseInfoModel, error)
}
