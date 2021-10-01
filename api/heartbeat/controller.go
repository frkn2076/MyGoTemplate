package heartbeat

import (
	"strconv"

	"app/MyGoTemplate/infra/cache"

	"github.com/gin-gonic/gin"
)

func NewController() *Controller {
	return &Controller{}
}

type Controller struct{}

func (u *Controller) GetAllReports(c *gin.Context) {
	c.JSON(200,
		"Cache average access time: "+strconv.FormatInt(cache.GetAvaregeAccessTime(), 10)+" milliseconds\n",
	)
}

func (u *Controller) ClearCache(c *gin.Context) {
	cache.Reset()
	c.AbortWithStatus(200)
}
