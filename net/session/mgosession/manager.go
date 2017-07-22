package mgosession

import (
	"github.com/caowenhua/cframe/net/session"
	"net/http"
)

type SessionManager struct {
	maxAge     int64
	cookieName string
	path       string
	domain     string
}

func (sm *SessionManager) newCookie() *http.Cookie {
	return &http.Cookie{
		Name:   sm.cookieName,
		Domain: sm.domain,
		Path:   sm.path,
		MaxAge: sm.maxAge,
	}
}

func (sm *SessionManager) NewSession(w http.ResponseWriter, r *http.Request) error {
	cookie := sm.newCookie()
	http.SetCookie(w, cookie)
}

func (sm *SessionManager) DestroySession(sid string) error {

}

func (sm *SessionManager) FindSession(sid string) (*session.Session, error) {

}
