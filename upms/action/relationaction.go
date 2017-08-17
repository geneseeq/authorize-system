package action

import (
	"sync"

	"github.com/geneseeq/authorize-system/upms/user"
	"github.com/geneseeq/authorize-system/db"
	"gopkg.in/mgo.v2/bson"
)

////////////////////////////////////////////
//user & role relation
////////////////////////////////////////////

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
			if v != "" {
				tmpDict[v] = "true"
			}
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

////////////////////////////////////////////
//group & role relation
////////////////////////////////////////////

type groupRoleRelationRoleRepository struct {
	mtx        sync.RWMutex
	collection string
	db         string
}

func (r *groupRoleRelationRoleRepository) FindFromGroup(id string) (*user.GroupRelationModel, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.GroupRelationModel{}
	err := con.Find(bson.M{"groupid": id}).One(&result)
	if err != nil {
		return nil, user.ErrUnknown
	}
	return &result, nil
}

func (r *groupRoleRelationRoleRepository) FindAllFromGroup() []*user.GroupRelationModel {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	var result []*user.GroupRelationModel
	con.Find(nil).All(&result)
	return result
}

func (r *groupRoleRelationRoleRepository) Store(g *user.GroupRelationModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	err := con.Insert(g)
	return err
}

func (r *groupRoleRelationRoleRepository) Remove(group_id string, g *user.GroupRelationModel) error {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.GroupRelationModel{}
	err := con.Find(bson.M{"groupid": group_id}).One(&result)
	if err != nil {
		return user.ErrUnknown
	}
	var newRoleID []string
	if result.RoleID != nil {
		tmpDict := map[string]string{}
		for _, v := range result.RoleID {
			tmpDict[v] = "true"
		}
		for _, v := range g.RoleID {
			delete(tmpDict, v)
		}
		for key, _ := range tmpDict {
			newRoleID = append(newRoleID, key)
		}
		result.RoleID = newRoleID
	}
	err = con.Update(bson.M{"groupid": group_id}, result)
	if err != nil {
		return user.ErrUnknown
	}
	return nil
}

func (r *groupRoleRelationRoleRepository) Update(id string, g *user.GroupRelationModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.GroupRelationModel{}
	err := con.Find(bson.M{"groupid": g.GroupID}).One(&result)
	if err != nil {
		return user.ErrUnknown
	}
	var newRoleID []string
	if result.RoleID != nil {
		result.RoleID = append(result.RoleID, g.RoleID...)
		tmpDict := map[string]string{}
		for _, v := range result.RoleID {
			if v != "" {
				tmpDict[v] = "true"
			}
		}
		for key, _ := range tmpDict {
			newRoleID = append(newRoleID, key)
		}
		result.RoleID = newRoleID
	}
	err = con.Update(bson.M{"groupid": g.GroupID}, result)
	if err != nil {
		return user.ErrUnknown
	}
	return err
}

// NewUserRelationRoleRepository returns a new instance of a in-memory cargo repository.
func NewGroupRoleRelationRoleRepository(db string, collection string) user.GroupRelationRepository {
	return &groupRoleRelationRoleRepository{
		db:         db,
		collection: collection,
	}
}

////////////////////////////////////////////
//group & user relation
////////////////////////////////////////////

type groupUserRelationRoleRepository struct {
	mtx        sync.RWMutex
	collection string
	db         string
}

func (r *groupUserRelationRoleRepository) FindFromGroup(id string) (*user.GroupRelationModel, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.GroupRelationModel{}
	err := con.Find(bson.M{"groupid": id}).One(&result)
	if err != nil {
		return nil, user.ErrUnknown
	}
	return &result, nil
}

func (r *groupUserRelationRoleRepository) FindAllFromGroup() []*user.GroupRelationModel {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	var result []*user.GroupRelationModel
	con.Find(nil).All(&result)
	return result
}

func (r *groupUserRelationRoleRepository) Store(g *user.GroupRelationModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	err := con.Insert(g)
	return err
}

func (r *groupUserRelationRoleRepository) Remove(group_id string, g *user.GroupRelationModel) error {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.GroupRelationModel{}
	err := con.Find(bson.M{"groupid": group_id}).One(&result)
	if err != nil {
		return user.ErrUnknown
	}
	var newUserID []string
	if result.UserID != nil {
		tmpDict := map[string]string{}
		for _, v := range result.UserID {
			tmpDict[v] = "true"
		}
		for _, v := range g.UserID {
			delete(tmpDict, v)
		}
		for key, _ := range tmpDict {
			newUserID = append(newUserID, key)
		}
		result.UserID = newUserID
	}
	err = con.Update(bson.M{"groupid": group_id}, result)
	if err != nil {
		return user.ErrUnknown
	}
	return nil
}

func (r *groupUserRelationRoleRepository) Update(id string, g *user.GroupRelationModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.GroupRelationModel{}
	err := con.Find(bson.M{"groupid": g.GroupID}).One(&result)
	if err != nil {
		return user.ErrUnknown
	}
	var newUserID []string
	if result.UserID != nil {
		result.UserID = append(result.UserID, g.UserID...)
		tmpDict := map[string]string{}
		for _, v := range result.UserID {
			if v != "" {
				tmpDict[v] = "true"
			}
		}
		for key, _ := range tmpDict {
			newUserID = append(newUserID, key)
		}
		result.UserID = newUserID
	}
	err = con.Update(bson.M{"groupid": g.GroupID}, result)
	if err != nil {
		return user.ErrUnknown
	}
	return err
}

// NewUserRelationRoleRepository returns a new instance of a in-memory cargo repository.
func NewGroupUserRelationRoleRepository(db string, collection string) user.GroupRelationRepository {
	return &groupUserRelationRoleRepository{
		db:         db,
		collection: collection,
	}
}

////////////////////////////////////////////
//role & authority relation
////////////////////////////////////////////

type roleAuthorityRelationRepository struct {
	mtx        sync.RWMutex
	collection string
	db         string
}

func (r *roleAuthorityRelationRepository) FindFromAuthority(id string) (*user.AuthorityRelationModel, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.AuthorityRelationModel{}
	err := con.Find(bson.M{"roleid": id}).One(&result)
	if err != nil {
		return nil, user.ErrUnknown
	}
	return &result, nil
}

func (r *roleAuthorityRelationRepository) FindAllFromAuthority() []*user.AuthorityRelationModel {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	var result []*user.AuthorityRelationModel
	con.Find(nil).All(&result)
	return result
}

func (r *roleAuthorityRelationRepository) Store(g *user.AuthorityRelationModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	err := con.Insert(g)
	return err
}

func (r *roleAuthorityRelationRepository) Remove(roleID string, d *user.DeleteAuthorityModel) error {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.AuthorityRelationModel{}
	err := con.Find(bson.M{"roleid": roleID}).One(&result)
	if err != nil {
		return user.ErrUnknown
	}
	var newAuthority []user.AuthorityModel
	if result.Authority != nil {
		tmpDict := map[string]user.AuthorityModel{}
		for _, v := range result.Authority {
			tmpDict[v.DataID] = v
		}
		for _, v := range d.DataID {
			delete(tmpDict, v)
		}
		for _, value := range tmpDict {
			newAuthority = append(newAuthority, value)
		}
		result.Authority = newAuthority
	}
	err = con.Update(bson.M{"roleid": roleID}, result)
	if err != nil {
		return user.ErrUnknown
	}
	return nil
}

func (r *roleAuthorityRelationRepository) Update(id string, a *user.AuthorityRelationModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.AuthorityRelationModel{}
	err := con.Find(bson.M{"roleid": id}).One(&result)
	if err != nil {
		return user.ErrUnknown
	}
	var newAuthority []user.AuthorityModel
	if result.Authority != nil {
		result.Authority = append(result.Authority, a.Authority...)
		tmpDict := map[string]user.AuthorityModel{}
		for _, v := range result.Authority {
			tmpDict[v.DataID] = v
		}
		for _, value := range tmpDict {
			newAuthority = append(newAuthority, value)
		}
		result.Authority = newAuthority
	}
	err = con.Update(bson.M{"roleid": id}, result)
	if err != nil {
		return user.ErrUnknown
	}
	return err
}

// NewRoleAuthorityRelationRepository returns a new instance of a in-memory cargo repository.
func NewRoleAuthorityRelationRepository(db string, collection string) user.AuthorityRelationRepository {
	return &roleAuthorityRelationRepository{
		db:         db,
		collection: collection,
	}
}
