package user

import (
	"time"
)

// DataSetModel is group struct
type DataSetModel struct {
	ID           string
	Rule         string
	Name         string
	MatchField   []MatchField
	Type         string
	Validity     bool
	Buildin      bool
	CreateUserID string
	CreateTime   time.Time
}

// MatchField is group struct
type MatchField struct {
	DataType  string   `json:"data_type"`
	SrcField  []string `json:"src_field"`
	DestField string   `json:"dest_field"`
}

// NewDataSet is create instance
func NewDataSet(id string) *DataSetModel {
	return &DataSetModel{
		ID: id,
	}
}

// DataSetRepository is user interface
type DataSetRepository interface {
	Store(set *DataSetModel) error
	FindDataSet(id string) (*DataSetModel, error)
	FindAllDataSet() []*DataSetModel
	Remove(id string) error
	Update(id string, set *DataSetModel) error
}
