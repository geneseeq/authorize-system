package user

import (
	"time"
)

// RoleDistributeModel is group struct
type RoleDistributeModel struct {
	UnionID      string    `bson:"_id"` //唯一ID
	ID           string    `bson:"id"`
	GroupID      string    `bson:"group_id"` //组类型（所有具备父子关系的组：机构，部门，项目组）
	UserID       string    `bson:"user_id"`  //组织编码
	RoleID       string    `bson:"role_id"`
	Buildin      bool      `bson:"buildin"`        //是否内建，true内建，false非内建
	CreateUserID string    `bson:"create_user_id"` //创建人ID
	UpdateUserID string    `bson:"update_user_id"`
	CreateTime   time.Time `bson:"create_time"` //创建时间
	UpdateTime   time.Time `bson:"update_time"` //更新时间
}

// NewRoleDistribute is create instance
func NewRoleDistribute(id string) *RoleDistributeModel {
	return &RoleDistributeModel{
		UnionID: id,
	}
}

// RoleDistributeRepository is user interface
type RoleDistributeRepository interface {
	Store(roleDistribute *RoleDistributeModel) error
	Find(id string) (*RoleDistributeModel, error)
	FindRoleDistributeAll() []*RoleDistributeModel
	Remove(id string) error
	Update(id string, roleDistribute *RoleDistributeModel) error
}
