package action

import (
	"sync"

	"github.com/geneseeq/authorize-system/cms/user"
	"github.com/geneseeq/authorize-system/db"
	"gopkg.in/mgo.v2/bson"
)

type userRelationRoleRepository struct {
	mtx        sync.RWMutex
	collection string
	db         string
}

func (r *userRelationRoleRepository) FindFromUser(id string) (*user.RoleRelationModel, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.RoleRelationModel{}
	err := con.Find(bson.M{"userid": id}).One(&result)
	if err != nil {
		return nil, user.ErrUnknown
	}
	return &result, nil
}

func (r *userRelationRoleRepository) FindAllFromUser() []*user.RoleRelationModel {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	var result []*user.RoleRelationModel
	con.Find(nil).All(&result)
	return result
}

func (r *userRelationRoleRepository) Store(g *user.RoleRelationModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	err := con.Insert(g)
	return err
}

func (r *userRelationRoleRepository) Remove(user_id string, role_id []string) error {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.RoleRelationModel{}
	err := con.Find(bson.M{"userid": user_id}).One(&result)
	if err != nil {
		return user.ErrUnknown
	}
	var newRoleID []string
	if result.RoleID != nil {
		tmpDict := map[string]string{}
		for _, v := range result.RoleID {
			tmpDict[v] = "true"
		}
		for _, v := range role_id {
			delete(tmpDict, v)
		}
		for key, _ := range tmpDict {
			newRoleID = append(newRoleID, key)
		}
		result.RoleID = newRoleID
	}
	err = con.Update(bson.M{"userid": user_id}, result)
	if err != nil {
		return user.ErrUnknown
	}
	return nil
}

func (r *userRelationRoleRepository) Update(id string, g *user.RoleRelationModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.RoleRelationModel{}
	err := con.Find(bson.M{"userid": g.UserID}).One(&result)
	if err != nil {
		return user.ErrUnknown
	}
	var newRoleID []string
	if result.RoleID != nil {
		result.RoleID = append(result.RoleID, g.RoleID...)
		tmpDict := map[string]string{}
		for _, v := range result.RoleID {
			tmpDict[v] = "true"
		}
		for key, _ := range tmpDict {
			newRoleID = append(newRoleID, key)
		}
		result.RoleID = newRoleID
	}
	err = con.Update(bson.M{"userid": g.UserID}, result)
	if err != nil {
		return user.ErrUnknown
	}
	return err
}

// NewUserRelationRoleRepository returns a new instance of a in-memory cargo repository.
func NewUserRelationRoleRepository(db string, collection string) user.RelationRepository {
	return &userRelationRoleRepository{
		db:         db,
		collection: collection,
	}
}
