package controllers

import(
	"strconv"

	"app/MyGoTemplate/cache"

	"github.com/gin-gonic/gin"
)

type ReportController struct{}

func (u *ReportController) GetAllReports(c *gin.Context) {
	
	c.JSON(200, 
		"Cache average access time: " + strconv.FormatInt(cache.GetAvaregeAccessTime(), 10) + " milliseconds\n",
		
	)
}


