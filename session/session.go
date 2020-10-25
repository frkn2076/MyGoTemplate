package session

import (
	"fmt"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"app/MyGoTemplate/logger"

	"github.com/gin-gonic/gin"
)

// Flash messages are session values that last until read
// When we request a flash message, it is removed from the session. To add a flash,
// call session.AddFlash(), and to get all flashes, call session.Flashes().

// store will hold all session data
var store *sessions.CookieStore

func SessionGet(c *gin.Context, key string) (interface{}, error)  {
	session, err := store.Get(c.Request, "cookie")
	if err != nil {
		logger.ErrorLog("An error occured while session get - SessionGet ", err.Error())
		return nil, err
	}
	return session.Values[key], nil
}

func SessionSet(c *gin.Context, key string, value interface{}, preparedSession *sessions.Session) (*sessions.Session, error) {
	if preparedSession != nil {
		preparedSession.Values[key] = value
		return preparedSession, nil
	} else {
		session, err := store.Get(c.Request, "cookie")
		if err != nil {
			logger.ErrorLog("An error occured while session set - SessionSet ", err.Error())
			return nil, err
		}
		session.Values[key] = value
		return session, nil
	}
}

func SessionAddFlash(c *gin.Context, key string, value string, preparedSession *sessions.Session) (*sessions.Session, error) {
	if preparedSession != nil {
		preparedSession.AddFlash(key, value)
		return preparedSession, nil
	} else {
		session, err := store.Get(c.Request, "cookie")
		if err != nil {
			logger.ErrorLog("An error occured while session get - SessionGet ", err.Error())
			return nil, err
		}
		return session, nil
	}
}

func SessionGetFlash(c *gin.Context, key string) (string, error) {
	session, err := store.Get(c.Request, "cookie")
	if err != nil {
		logger.ErrorLog("An error occured while session get - SessionGet ", err.Error())
		return "", err
	}
	value := session.Flashes(key)
	return fmt.Sprintf("%v", value), nil
}

func SessionSave(c *gin.Context, s *sessions.Session) error {
	if err := s.Save(c.Request, c.Writer); err != nil {
		logger.ErrorLog("An error occured while session save - SessionSet", err.Error())
		return err
	}
	return nil
}

func init() {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	cookieStore := sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)

	cookieStore.Options = &sessions.Options{
		MaxAge:   60 * 10, //10 minutes
		HttpOnly: true,
	}

	store = cookieStore
}

//#region Helper



//#endregion