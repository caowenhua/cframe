package sequence

import (
	"errors"
	"github.com/caowenhua/cframe/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const CN_SEQ = "mgo_sequence"

type MgoSequence struct {
	DB         *mgo.Database
	Collection string
}

func NewMgoSequence(db *mgo.Database) (*MgoSequence, error) {
	return NewMgoSequence2(db, CN_SEQ)
}

func NewMgoSequence2(db *mgo.Database, collectionName string) (*MgoSequence, error) {
	if db == nil {
		return nil, errors.New("database is null")
	}
	if collectionName == "" {
		return nil, errors.New("collection name is null")
	}
	return &MgoSequence{
		DB:         db,
		Collection: collectionName,
	}, nil
}

//query one sequence by id
func (mg *MgoSequence) Query(id string) (uint64, error) {
	var seq = util.Map{}
	_, err := mg.DB.C(mg.Collection).Find(bson.M{"_id": id}).Apply(
		mgo.Change{
			Update:    bson.M{"$inc": bson.M{"val": 1}},
			Upsert:    true,
			ReturnNew: true,
		}, &seq,
	)
	return seq.GetUint("val"), err
}
