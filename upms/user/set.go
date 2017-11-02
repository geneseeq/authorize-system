package user

import (
	"time"
)

// DataSetModel is group struct
type DataSetModel struct {
	UnionID      string       `bson:"_id"`
	ID           string       `bson:"id"`
	Rule         string       `bson:"rule"`
	Name         string       `bson:"name"`
	MatchField   []MatchField `bson:"match_field"`
	Type         string       `bson:"type"`
	Validity     bool         `bson:"validity"`
	Buildin      bool         `bson:"buildin"`
	CreateUserID string       `bson:"create_user_id"`
	UpdateUserID string       `bson:"update_user_id"`
	CreateTime   time.Time    `bson:"create_time"`
	UpdateTime   time.Time    `bson:"update_time"`
}

// MatchField 解决医学部访问他人数据
type MatchField struct {
	DataType  int      `json:"data_type" bson:"data_type"`   //数据类型：0表示个人数据，1表示组数据
	SrcField  []string `json:"src_field" bson:"src_field"`   //表示多个字段，譬如多个组id，多个个人id
	DestField string   `json:"dest_field" bson:"dest_field"` //基础数据表中字段
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
