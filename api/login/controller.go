package login

import (
	"math/rand"
	"os"

	"app/MyGoTemplate/infra/customeerror"
	"app/MyGoTemplate/api"
	"app/MyGoTemplate/infra/db"
	"app/MyGoTemplate/infra/db/login"
	"app/MyGoTemplate/infra/helper"
	"app/MyGoTemplate/infra/logger"
	"app/MyGoTemplate/infra/resource"
	s "app/MyGoTemplate/infra/session"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func NewController() *Controller {
	return &Controller{LoginRepository: login.NewRepository()}
}

type Controller struct{
	LoginRepository *login.Repository
}

func (u *Controller) Register(c *gin.Context) {

	session, err := s.Store.Get(c.Request, os.Getenv("SessionCookieName"))
	if err != nil {
		logger.ErrorLog("An error occured while session get - Register - login.go", err.Error())
		c.AbortWithStatus(500)
		return
	}

	session.Values["language"] = "TR"

	// session, _ := s.SessionSet(c, "version", "1.0.0", nil)

	var loginRequest LoginRequest
	if err := c.Bind(&loginRequest); err != nil {
		logger.ErrorLog("Invalid request - Register - login.go", err.Error())
		c.AbortWithStatus(500)
		return
	}

	if !helper.IsEmailValid(loginRequest.Email) {
		logger.ErrorLog("Invalid mail", loginRequest.Email)
		c.Error(customeerror.NotAValidEmail)
		return
	}

	if len(loginRequest.Password) < 6 {
		logger.ErrorLog("Short password - Login - login.go")
		c.Error(customeerror.ShortPassword)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(loginRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.ErrorLog("An error occured while generating crypted password - Register - login.go", err.Error())
		c.AbortWithStatus(500)
		return
	}

	login := login.Entity{UserName: loginRequest.UserName, Email: loginRequest.Email, Password: string(hashedPassword)}

	isMailValidationActive := resource.GetValue("isMailValidatonActive") == "true"

	if isMailValidationActive {
		key := rand.Intn(899999) + 100000
		
		session.AddFlash("name", login.UserName) 
		session.AddFlash("email", login.Email) 
		session.AddFlash("hashedPassword", string(hashedPassword)) 
		session.AddFlash("emailKey", string(key))

		helper.SendMail("ozturkfurkan1994@hotmail.com", string(key), session.Values["language"].(string))
	} else {
		tx := db.GormDB.Begin()

		if err := u.LoginRepository.Create(tx, login); err != nil {
			c.Error(customeerror.UserAlreadyExists)
			return
		}

		tx.Commit()

		session.Values["isActive"] = true
	}

	if err := session.Save(c.Request, c.Writer); err != nil {
		logger.ErrorLog("An error occured while saving session - Register - login.go", err.Error())
		c.AbortWithStatus(500)
		return
	}

	c.JSON(200, api.Success)
}

func (u *Controller) RegisterValidation(c *gin.Context) {
	session, err := s.Store.Get(c.Request, os.Getenv("SessionCookieName"))
	if err != nil {
		logger.ErrorLog("An error occured while session get - Register - login.go", err.Error())
		c.AbortWithStatus(500)
		return
	}

	var registerValidationRequest RegisterValidationRequest
	if err := c.Bind(&registerValidationRequest); err != nil {
		logger.ErrorLog("Invalid request - RegisterValidation - login.go", err.Error())
		c.AbortWithStatus(500)
		return
	}

	emailKey := session.Flashes("emailKey")[0]
	
	if emailKey == registerValidationRequest.Key {
		login := login.Entity{UserName: session.Flashes("name")[0].(string), Email: session.Flashes("email")[0].(string), Password: session.Flashes("hashedPassword")[0].(string)}

		tx := db.GormDB.Begin()

		if err := u.LoginRepository.Create(tx, login); err != nil {
			session.AddFlash("name", login.UserName) 
			session.AddFlash("email", login.Email) 
			session.AddFlash("hashedPassword", login.Password)  
			session.AddFlash("emailKey", emailKey.(string))
			c.Error(customeerror.UserAlreadyExists)
			return
		}

		tx.Commit()

		session.Values["isActive"] = true

		c.JSON(200, api.Success)
	}
}

func (u *Controller) Login(c *gin.Context) {

	session, err := s.Store.Get(c.Request, os.Getenv("SessionCookieName"))
	if err != nil {
		logger.ErrorLog("An error occured while session get - Register - login.go", err.Error())
		c.AbortWithStatus(500)
		return
	}

	var loginRequest LoginRequest
	if err := c.Bind(&loginRequest); err != nil {
		logger.ErrorLog("Invalid request - Login - login.go", err.Error())
		c.AbortWithStatus(500)
		return
	}

	tx := db.GormDB.Begin()

	login, err := u.LoginRepository.First(tx, loginRequest.UserName)
	if err != nil {
		c.Error(customeerror.UserNotFound)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(loginRequest.Password)); err != nil {
		logger.ErrorLog(err)
		c.Error(customeerror.WrongPassword)
		return
	}

	tx.Commit()

	session.Values["isActive"] = true

	session.Save(c.Request, c.Writer)

	c.JSON(200, gin.H{
		"isSuccess": true,
	})
}
