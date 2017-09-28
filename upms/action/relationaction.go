package action

import (
	"sync"
	"time"

	"github.com/geneseeq/authorize-system/db"
	"github.com/geneseeq/authorize-system/upms/user"
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
	err := con.Find(bson.M{"user_id": id}).One(&result)
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

func (r *userRelationRoleRepository) Remove(userID string, role *user.RoleRelationModel) error {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.RoleRelationModel{}
	err := con.Find(bson.M{"user_id": userID}).One(&result)
	if err != nil {
		return user.ErrUnknown
	}
	if len(role.RoleID) == 0 {
		if len(result.RoleID) == 0 {
			err = con.Remove(bson.M{"user_id": userID})
		} else {
			return user.ErrRemove
		}
		if err != nil {
			return user.ErrUnknown
		}

	} else {
		var newRoleID []string
		if result.RoleID != nil {
			tmpDict := map[string]string{}
			for _, v := range result.RoleID {
				tmpDict[v] = "true"
			}
			for _, v := range role.RoleID {
				delete(tmpDict, v)
			}
			for key, _ := range tmpDict {
				newRoleID = append(newRoleID, key)
			}
			result.RoleID = newRoleID
		}
		result.UpdateTime = user.TimeUtcToCst(time.Now())
		err = con.Update(bson.M{"user_id": userID}, result)
		if err != nil {
			return user.ErrUnknown
		}
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
	err := con.Find(bson.M{"user_id": g.UserID}).One(&result)
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
	result.UpdateTime = user.TimeUtcToCst(time.Now())
	if g.CreateUserID != "" {
		result.CreateUserID = g.CreateUserID
	}
	if g.Buildin != false {
		result.Buildin = g.Buildin
	}
	err = con.Update(bson.M{"user_id": g.UserID}, result)
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
	err := con.Find(bson.M{"group_id": id}).One(&result)
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

func (r *groupRoleRelationRoleRepository) Remove(groupID string, g *user.GroupRelationModel) error {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.GroupRelationModel{}
	err := con.Find(bson.M{"group_id": groupID}).One(&result)
	if err != nil {
		return user.ErrUnknown
	}
	if len(g.UserID) == 0 && len(g.RoleID) == 0 {
		if len(result.UserID) == 0 && len(result.RoleID) == 0 {
			err = con.Remove(bson.M{"group_id": groupID})
		} else {
			return user.ErrRemove
		}
		if err != nil {
			return user.ErrUnknown
		}
	} else {

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
		result.UpdateTime = user.TimeUtcToCst(time.Now())
		err = con.Update(bson.M{"group_id": groupID}, result)
		if err != nil {
			return user.ErrUnknown
		}
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
	err := con.Find(bson.M{"group_id": g.GroupID}).One(&result)
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
	result.UpdateTime = user.TimeUtcToCst(time.Now())
	if g.CreateUserID != "" {
		result.CreateUserID = g.CreateUserID
	}
	if g.Buildin != false {
		result.Buildin = g.Buildin
	}
	err = con.Update(bson.M{"group_id": g.GroupID}, result)
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
	err := con.Find(bson.M{"group_id": id}).One(&result)
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

func (r *groupUserRelationRoleRepository) Remove(groupID string, g *user.GroupRelationModel) error {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := user.GroupRelationModel{}
	err := con.Find(bson.M{"group_id": groupID}).One(&result)
	if err != nil {
		return user.ErrUnknown
	}
	if len(g.UserID) == 0 && len(g.RoleID) == 0 {
		if len(result.UserID) == 0 && len(result.RoleID) == 0 {
			err = con.Remove(bson.M{"group_id": groupID})
		} else {
			return user.ErrRemove
		}
		if err != nil {
			return user.ErrUnknown
		}

	} else {
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
		result.UpdateTime = user.TimeUtcToCst(time.Now())
		err = con.Update(bson.M{"group_id": groupID}, result)
		if err != nil {
			return user.ErrUnknown
		}
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
	err := con.Find(bson.M{"group_id": g.GroupID}).One(&result)
	if err != nil {
		return user.ErrUnknown
	}
	var newUserID []string
	if result.UserID != nil {
		result.UserID = append(result.UserID, g.UserID...)
		//去重
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
	result.UpdateTime = user.TimeUtcToCst(time.Now())
	if g.CreateUserID != "" {
		result.CreateUserID = g.CreateUserID
	}
	if g.Buildin != false {
		result.Buildin = g.Buildin
	}
	err = con.Update(bson.M{"group_id": g.GroupID}, result)
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
	err := con.Find(bson.M{"role_id": id}).One(&result)
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
	err := con.Find(bson.M{"role_id": roleID}).One(&result)
	if err != nil {
		return user.ErrUnknown
	}
	if len(d.DataID) == 0 {
		if len(result.Authority) == 0 {
			err = con.Remove(bson.M{"role_id": roleID})
		} else {
			return user.ErrRemove
		}
		if err != nil {
			return user.ErrUnknown
		}

	} else {
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
		err = con.Update(bson.M{"role_id": roleID}, result)
		if err != nil {
			return user.ErrUnknown
		}
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
	err := con.Find(bson.M{"role_id": id}).One(&result)
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
	err = con.Update(bson.M{"role_id": id}, result)
	if a.CreateUserID != "" {
		result.CreateUserID = a.CreateUserID
	}
	if a.Buildin != false {
		result.Buildin = a.Buildin
	}
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
