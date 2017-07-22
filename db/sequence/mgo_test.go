package sequence

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestMgoSequence(t *testing.T) {
	session, err := mgo.Dial("cwh:123@127.0.0.1:27017/test")
	if err != nil {
		t.Error(err)
		return
	}
	db := session.DB("test")

	{
		//default
		defer func() {
			db.C(CN_SEQ).RemoveAll(nil)
		}()

		mgoSeq, err := NewMgoSequence(db)
		if err != nil {
			t.Error(err)
			return
		}

		var val uint64 = 0
		var id = bson.NewObjectId().Hex()
		for i := 0; i < 100; i++ {
			v, err := mgoSeq.Query(id)
			if err != nil {
				t.Error(err)
				return
			}
			if v != val+1 {
				t.Error("sequence err", v, val, i)
				return
			}
			val = v
		}
	}

	{
		//diy collection name
		defer func() {
			db.C(CN_SEQ + "_test").RemoveAll(nil)
		}()

		mgoSeq, err := NewMgoSequence2(db, CN_SEQ+"_test")
		if err != nil {
			t.Error(err)
			return
		}

		var val uint64 = 0
		var id = bson.NewObjectId().Hex()
		for i := 0; i < 100; i++ {
			v, err := mgoSeq.Query(id)
			if err != nil {
				t.Error(err)
				return
			}
			if v != val+1 {
				t.Error("sequence err", v, val, i)
				return
			}
			val = v
		}

		//test error
		_, err = NewMgoSequence2(nil, CN_SEQ)
		if err == nil {
			t.Error("db is null but no err")
			return
		}

		_, err = NewMgoSequence2(db, "")
		if err == nil {
			t.Error("name is null but no err")
			return
		}
	}
}
