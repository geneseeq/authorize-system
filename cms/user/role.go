package user

// GroupModel is group struct
type RoleModel struct {
	Parent         string
	ID             string
	Type           int
	Name           string
	Code           string
	Alias          string
	Buildin        bool
	Create_user_id string
	Create_time    string
}

// NewRole is create instance
func NewRole(id string) *RoleModel {
	return &RoleModel{
		ID: id,
	}
}

// RoleRepository is user interface
type RoleRepository interface {
	Store(role *RoleModel) error
	Find(id string) (*RoleModel, error)
	FindRoleAll() []*RoleModel
	Remove(id string) error
	Update(id string, role *RoleModel) error
}
