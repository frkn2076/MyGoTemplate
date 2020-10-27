package controllers

import (
	"os"

	"app/MyGoTemplate/controllers/models/request"
	"app/MyGoTemplate/controllers/models/response"
	"app/MyGoTemplate/db"
	"app/MyGoTemplate/db/entities"
	"app/MyGoTemplate/db/repo"
	s "app/MyGoTemplate/session"
	"app/MyGoTemplate/logger"
	"app/MyGoTemplate/definedErrors"
	"app/MyGoTemplate/helper"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct{}

func (u *LoginController) Register(c *gin.Context) {

	session, err := s.Store.Get(c.Request, os.Getenv("SessionCookieName"))
	if err != nil {
		logger.ErrorLog("An error occured while session get - Register - login.go", err.Error())
		c.AbortWithStatus(500)
		return
	}

	session.Values["language"] = "TR"

	// session, _ := s.SessionSet(c, "version", "1.0.0", nil)

	var loginRequest request.LoginRequest
	if err := c.Bind(&loginRequest); err != nil {
		logger.ErrorLog("Invalid request - Register - login.go", err.Error())
		c.AbortWithStatus(500)
		return
	}

	if !helper.IsEmailValid(loginRequest.Email) {
		logger.ErrorLog("Invalid mail", loginRequest.Email)
		c.Error(definedErrors.NotAValidEmail)
		return
	}

	if len(loginRequest.Password) < 6 {
		logger.ErrorLog("Short password - Login - login.go")
		c.Error(definedErrors.ShortPassword)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(loginRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.ErrorLog("An error occured while generating crypted password - Register - login.go", err.Error())
		c.AbortWithStatus(500)
		return
	}

	tx := db.GormDB.Begin()

	login := entities.Login{UserName: loginRequest.UserName, Email: loginRequest.Email, Password: string(hashedPassword)}

	if err := repo.Login.Create(tx, login); err != nil {
		c.Error(definedErrors.UserAlreadyExists)
		return
	}

	tx.Commit()

	session.Values["isActive"] = true

	if err := session.Save(c.Request, c.Writer); err != nil {
		logger.ErrorLog("An error occured while saving session - Register - login.go", err.Error())
		c.AbortWithStatus(500)
		return
	}
	
	c.JSON(200, response.Success)
}

func (u *LoginController) Login(c *gin.Context) {

	session, err := s.Store.Get(c.Request, os.Getenv("SessionCookieName"))
	if err != nil {
		logger.ErrorLog("An error occured while session get - Register - login.go", err.Error())
		c.AbortWithStatus(500)
		return
	}

	var loginRequest request.LoginRequest
	if err := c.Bind(&loginRequest); err != nil {
		logger.ErrorLog("Invalid request - Login - login.go", err.Error())
		c.AbortWithStatus(500)
		return
	}

	tx := db.GormDB.Begin()

	login, err := repo.Login.First(tx, loginRequest.UserName)
	if err != nil {
		c.Error(definedErrors.UserNotFound)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(loginRequest.Password)); err != nil {
		logger.ErrorLog(err)
		c.Error(definedErrors.WrongPassword)
		return
	}

	tx.Commit()

	session.Values["isActive"] = true

	session.Save(c.Request, c.Writer)

	c.JSON(200, gin.H{
		"isSuccess": true,
	})
}
