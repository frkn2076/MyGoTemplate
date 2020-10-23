package session

import (
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"app/MyGoTemplate/logger"

	"github.com/gin-gonic/gin"
)

// Flash messages are session values that last until read
// When we request a flash message, it is removed from the session. To add a flash,
// call session.AddFlash(), and to get all flashes, call session.Flashes().

// store will hold all session data
var store *sessions.CookieStore = initStore()

func SessionGet(c *gin.Context, key string) interface{} {
	session, err := store.Get(c.Request, "cookie-name")
	if err != nil {
		logger.ErrorLog("An error occured while session get - SessionGet ", err.Error())
	}
	return session.Values[key]
}

//Set session at once to provide Unit of Work
func SessionSet(c *gin.Context, key string, value interface{}){
	session, err := store.Get(c.Request, "cookie-name")
	if err != nil {
		logger.ErrorLog("An error occured while session set - SessionSet ", err.Error())
	}
	session.Values[key] = value

	err = session.Save(c.Request, c.Writer)
	if err != nil {
		logger.ErrorLog("An error occured while session save - SessionSet", err.Error())
	}
}


//#region Helper

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

//#endregion