package controllers

import(
	"strconv"

	"app/MyGoTemplate/cache"

	"github.com/gin-gonic/gin"
)

type ManagerController struct{}

func (u *ManagerController) GetAllReports(c *gin.Context) {
	
	c.JSON(200, 
		"Cache average access time: " + strconv.FormatInt(cache.GetAvaregeAccessTime(), 10) + " milliseconds\n",
		
	)
}

func (u *ManagerController) ClearCache(c *gin.Context) {
	cache.Reset()
	c.AbortWithStatus(200)
}


