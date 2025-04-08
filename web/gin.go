package web

import (
	"github.com/gin-gonic/gin"
	"miaoshaSystem/Middleware"
	"miaoshaSystem/user"
)

func Gin() {
	r := gin.Default()
	r.Use(Middleware.RateLimitMiddleware(10))
	r.POST("/user/register", user.Register)
	r.POST("/user/login", user.Login)
	r.POST("/createmiaosha", user.Createmiaosha) //因为这次考核的核心是秒杀系统，所以干脆把创建秒杀活动的代码和上架商品的代码写在一起了
	r.PUT("/miaosha/:productName", user.Miaosha)
	r.Run(":8080")

}
