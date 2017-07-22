package mgosession

import (
	"github.com/caowenhua/cframe/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const CN_SESSION = "mgo_session"

var db *mgo.Database

type MgoSession struct {
	Id   string   `bson:"_id" json:"_id"`
	Time int64    `bson:"time" json:"time"`
	Data util.Map `bson:"data" json:"data"`
}

func InitMgoSession(db *mgo.Database) {

}

func NewMgoSession() *MgoSession {
	return &MgoSession{
		Id:   bson.NewObjectId().Hex(),
		Time: util.Now(),
		Data: util.Map{},
	}
}

//set session value
func (ms *MgoSession) Set(value util.Map) error {
	err := db.C(CN_SESSION).Update(bson.M{"_id": ms.Id}, bson.M{"$set": value})
	return err
}

func (ms *MgoSession) Store() error {
	return db.C(CN_SESSION).Insert(ms)
}

//get session value
func (ms *MgoSession) Get(keys []string) (util.Map, error) {
	//var selector = bson.M{}
	//for _, key := range keys{
	//	selector["data."+key]=1
	//}
	//var session MgoSession
	//err := db.C(CN_SESSION).Find(bson.M{"_id":ms.Id}).Select(selector).One()
}

//delete session value
func (ms *MgoSession) Delete(keys []string) error {

}

//back current sessionID
func (ms *MgoSession) SessionID() string {

}
