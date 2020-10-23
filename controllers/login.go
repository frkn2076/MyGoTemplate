package controllers

import(
	"app/MyGoTemplate/db"
	"app/MyGoTemplate/controllers/models/request"
	"app/MyGoTemplate/db/entities"
	s "app/MyGoTemplate/session"
	"app/MyGoTemplate/cache"
	_ "app/MyGoTemplate/resource"

	"github.com/gin-gonic/gin"
)

type LoginController struct{}

func (u *LoginController) Login(c *gin.Context) {

	c.AbortWithStatus(500)
	return

	var loginRequest request.LoginRequest
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

	s.SessionSet(c, "isActive", true)
	s.SessionGet(c, "isActive")

	c.JSON(200, gin.H{
		"isSuccess": cache.Get("furkan") + cache.Get("cookie"),
	})
}

func (u *LoginController) Register(c *gin.Context) {
	// var login Login
	// err := json.NewDecoder(c.Body).Decode(&login)

	cache.Set("furkan","ozturk", -1)

	c.JSON(200, gin.H{
		"isSuccess": true,
	})
}