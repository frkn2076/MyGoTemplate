package controllers

import(
	"app/MyGoTemplate/db"
	"app/MyGoTemplate/controllers/models"
	"app/MyGoTemplate/db/entities"
	s "app/MyGoTemplate/session"

	"github.com/gin-gonic/gin"
)

type LoginController struct{}

func (u *LoginController) Login(c *gin.Context) {
	var loginRequest models.LoginRequest
	if err := c.Bind(&loginRequest); err != nil {
		c.JSON(400, err)
		return
	}

	login := entities.Login{UserName: loginRequest.UserName, Email: loginRequest.Email, Password: loginRequest.Password}

	if err := db.GormDB.Where("user_name = ? or email = ?", loginRequest.UserName, loginRequest.Email).First(&entities.Login{}).Error; err != nil {
		db.GormDB.Create(&login)
	} else {
		c.JSON(400, "This user already exists. Please check your User Name and Email")
		return
	}

	session, _ := s.Store.Get(c.Request, "cookie-name")

	session.Values["userName"] = loginRequest.UserName
	session.Values["authenticated"] = true

	_ = session.Save(c.Request, c.Writer)

	c.JSON(200, gin.H{
		"isSuccess": loginRequest.UserName,
	})
}

func (u *LoginController) Register(c *gin.Context) {
	// var login Login
	// err := json.NewDecoder(c.Body).Decode(&login)

	c.JSON(200, gin.H{
		"isSuccess": true,
	})
}