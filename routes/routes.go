package routes

import (
	"backend/controllers"
	"backend/logger"
	"backend/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true, // 允许携带 Cookie
	}))

	v1 := r.Group("/api")

	//登陆
	v1.GET("/nonce", controllers.GetNonceHandler)
	v1.POST("/login", controllers.LoginHandler)
	//检测用户是否登陆
	v1.GET("/check-login", controllers.CheckUserLoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware())
	{
		//用户部署合约
		v1.GET("/abi", controllers.GetAbiBytecodeHandler)
		v1.POST("/deploy", controllers.DeployHandler)
		//获取用户部署过的合约
		v1.GET("/get-token-list", controllers.GetTokenListHandler)
		//获取所有部署过的合约以及常见token的name和address
		v1.GET("/get-all-token", controllers.GetAllTokenAddressAndNameHandler)
		//上传合约媒体账号信息
		v1.POST("/post-token-info", controllers.PostTokenInfoHandler)
		//获取合约媒体账号信息
		v1.GET("/get-token-info", controllers.GetTokenInfoHandler)
		//添加consumer
		v1.GET("/add-consumer", controllers.AddConsumerHandler)
		//添加automation
		v1.GET("/add-automation", controllers.AddAutomationHandler)

	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "404",
		})
	})
	return r
}
