package router

import (
	"github.com/gin-gonic/gin"

	"redu/api/v1"
)

func SeckillRouter(appGroup *gin.RouterGroup) {
	v1SeckillGroup := appGroup.Group("v1/seckill_activity")
	v1SeckillGroup.POST("", v1.GroupApp.CreateSeckillActivity)
	v1SeckillGroup.GET("list", v1.GroupApp.GetSeckillList)
}
