package user

// GroupModel is group struct
type GroupModel struct {
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

// NewGroup is create instance
func NewGroup(id string) *GroupModel {
	return &GroupModel{
		ID: id,
	}
}

// GroupRepository is user interface
type GroupRepository interface {
	Store(group *GroupModel) error
	Find(id string) (*GroupModel, error)
	FindGroupAll() []*GroupModel
	Remove(id string) error
}
