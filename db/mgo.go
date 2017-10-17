package db

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	yaml "gopkg.in/yaml.v2"
)

type Mongo struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Db   string `yaml:"db"`
	Coll string `yaml:"coll"`
}

type DB struct {
	Mongo Mongo `yaml:"mongo"`
}

func getAddress(mongoCfg DB) (string, error) {
	address := strings.Join([]string{
		mongoCfg.Mongo.Host, ":", mongoCfg.Mongo.Port,
	}, "")
	return address, nil
}

var session *mgo.Session

const (
	defaultGOPATH = "C:/work/goworkspace"
)

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
func init() {

	goPath := envString("GOPATH", defaultGOPATH)
	confPath := goPath + "/src/github.com/geneseeq/authorize-system/conf/" + "conf.yaml"
	fmt.Println(confPath)
	content, _ := ioutil.ReadFile(confPath)
	mongoCfg := DB{}
	err := yaml.Unmarshal(content, &mongoCfg)
	address, err := getAddress(mongoCfg)
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{address},
		Direct:    false,
		Timeout:   time.Second * 1,
		PoolLimit: 4096, // Session.SetPoolLimit
	}
	session, err = mgo.DialWithInfo(dialInfo)

	if err != nil {
		log.Println(err.Error())
	}
	session.SetMode(mgo.Monotonic, true)

}

type SessionStore struct {
	session *mgo.Session
}

//获取数据库的collection
func (d *SessionStore) GetConnect(db string, collection string) *mgo.Collection {
	return d.session.DB(db).C(collection)
}

//为每一HTTP请求创建新的DataStore对象
func NewSessionStore() *SessionStore {
	ds := &SessionStore{
		session: session.Copy(),
	}
	return ds
}

func (d *SessionStore) Close() {
	d.session.Close()
}

func GetErrNotFound() error {
	return mgo.ErrNotFound
}
