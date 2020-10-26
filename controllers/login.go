package controllers

import (
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
)

type LoginController struct{}

func (u *LoginController) Register(c *gin.Context) {

	session, _ := s.SessionSet(c, "language", "TR", nil)
	// session, _ := s.SessionSet(c, "version", "1.0.0", nil)

	var loginRequest request.LoginRequest
	if err := c.Bind(&loginRequest); err != nil {
		logger.ErrorLog("Invalid request for login/login")
		c.AbortWithStatus(500)
		return
	}

	if !helper.IsEmailValid(loginRequest.Email) {
		logger.ErrorLog("Invalid mail")
		c.Error(definedErrors.NotAValidEmail)
		return
	}

	tx := db.GormDB.Begin()

	login := entities.Login{UserName: loginRequest.UserName, Email: loginRequest.Email, Password: loginRequest.Password}

	if err := repo.Login.Create(tx, login); err != nil {
		c.Error(definedErrors.UserAlreadyExists)
		return
	}

	tx.Commit()

	s.SessionSet(c, "isActive", true, session)

	s.SessionSave(c, session)

	c.JSON(200, response.Success)

}

func (u *LoginController) Login(c *gin.Context) {

	c.JSON(200, gin.H{
		"isSuccess": true,
	})
}
