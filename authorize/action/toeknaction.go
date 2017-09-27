package action

import (
	"sync"
	"time"

	"github.com/geneseeq/authorize-system/authorize/auth"
	"github.com/geneseeq/authorize-system/db"
	"gopkg.in/mgo.v2/bson"
)

type tokenDBRepository struct {
	mtx        sync.RWMutex
	collection string
	db         string
}

func (r *tokenDBRepository) Find(id string) (*auth.TokenModel, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := auth.TokenModel{}
	err := con.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return nil, auth.ErrUnknown
	}
	return &result, nil
}

func (r *tokenDBRepository) FindAllToken() []*auth.TokenModel {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	var result []*auth.TokenModel
	con.Find(nil).All(&result)
	return result
}

func (r *tokenDBRepository) Store(d *auth.TokenModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	err := con.Insert(d)
	return err
}

func (r *tokenDBRepository) Remove(id string) error {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	err := con.Remove(bson.M{"id": id})
	if err != nil {
		return auth.ErrUnknown
	}
	return nil
}

func newTokenModel(new *auth.TokenModel, result auth.TokenModel) auth.TokenModel {
	if new.AccessToken != "" {
		result.AccessToken = new.AccessToken
	}

	if new.Validity != false {
		result.Validity = new.Validity
	}

	if new.TokenType != "" {
		result.TokenType = new.TokenType
	}

	if new.ExpiresIn != 0 {
		result.ExpiresIn = new.ExpiresIn
	}

	if new.RefreshToken != "" {
		result.RefreshToken = new.RefreshToken
	}
	result.UpdateTime = auth.TimeUtcToCst(time.Now())
	return result
}

func (r *tokenDBRepository) Update(id string, d *auth.TokenModel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	ds := db.NewSessionStore()
	defer ds.Close()
	con := ds.GetConnect(r.db, r.collection)
	result := auth.TokenModel{}
	err := con.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return auth.ErrUnknown
	}
	result = newTokenModel(d, result)
	err = con.Update(bson.M{"id": id}, result)
	return err
}

// NewTokenDBRepository returns a new instance of a in-memory cargo repository.
func NewTokenDBRepository(db string, collection string) auth.TokenRepository {
	return &tokenDBRepository{
		db:         db,
		collection: collection,
	}
}
