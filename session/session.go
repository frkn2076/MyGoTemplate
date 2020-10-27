package session

import (
	"github.com/gorilla/sessions"
)

// store will hold all session data
var Store *sessions.CookieStore

func init() {

	cookieStore := sessions.NewCookieStore([]byte("secret-key"))

	cookieStore.Options = &sessions.Options{
		MaxAge:   60 * 10, //10 minutes
		HttpOnly: false,
	}

	Store = cookieStore
}
