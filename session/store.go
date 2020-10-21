package session

import(
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

type User struct {
	Username      string
	Authenticated bool
}	

// store will hold all session data
var Store *sessions.CookieStore = initStore()

func initStore() *sessions.CookieStore {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	store := sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)

	store.Options = &sessions.Options{
		MaxAge:   60 * 10, //10 minutes
		HttpOnly: true,
	}

	return store
}