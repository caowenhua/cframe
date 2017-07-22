package session

import "github.com/caowenhua/cframe/util"

type Manager interface {
	NewSession() error
	DestroySession(sid string) error
	FindSession(sid string) (*Session, error)
}

type Session interface {
	Set(value util.Map) error            //set session value
	Get(keys []string) (util.Map, error) //get session value
	Delete(keys []string) error          //delete session value
	SessionID() string                   //back current sessionID
}
