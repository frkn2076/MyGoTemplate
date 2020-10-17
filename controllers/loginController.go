package controllers

import(
	// "encoding/json"
	"app/MyGoTemplate/db"
	"app/MyGoTemplate/controllers/models"
	"app/MyGoTemplate/db/entities"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginRequest models.LoginRequest
	if err := c.Bind(&loginRequest); err != nil {
		c.JSON(400, err)
		return
	}

	login := entities.Login{UserName: loginRequest.UserName, Email: loginRequest.Email, Password: loginRequest.Password}

	if err := db.GormDB.Debug().Where("user_name = ? or email = ?", loginRequest.UserName, loginRequest.Email).First(&entities.Login{}).Error; err != nil {
		db.GormDB.Debug().Create(&login)
	} else {
		c.JSON(400, "This user already exists. Please check your User Name and Email")
	}

	c.JSON(200, gin.H{
		"isSuccess": loginRequest.UserName,
	})
}

func Register(c *gin.Context) {
	// var login Login
	// err := json.NewDecoder(c.Body).Decode(&login)

	c.JSON(200, gin.H{
		"isSuccess": true,
	})
}