package initialize

import (
	"MSC2021/src/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() (*gin.Engine, error) {
	router := gin.Default()

	apiRouter := router.Group("api")
	{
		accountRouter := apiRouter.Group("account")
		{
			accountRouter.POST("login")
			accountRouter.POST("register")
			accountRouter.Use(middleware.LoginRequired())
			accountRouter.GET("me")  // get profile
			accountRouter.PUT("me")  // change profile
			accountRouter.POST("me") // change password
			accountRouter.GET("logout")
		}
		apiRouter.GET("questions")
		questionRouter := apiRouter.Group("question")
		{
			questionRouter.GET(":id")
			questionRouter.Use(middleware.LoginRequired())
			questionRouter.POST(":id")
		}
		adminRouter := apiRouter.Group("admin")
		adminRouter.Use(middleware.LoginRequired())
		adminRouter.Use(middleware.AdminRequired())
		{
			adminRouter.GET("users")
			adminUserRouter := adminRouter.Group("user")
			{
				adminUserRouter.GET(":id")
				adminUserRouter.PUT(":id")
				adminUserRouter.DELETE(":id")
			}
			adminRouter.POST("create-user")
			adminRouter.GET("questions")
			adminQuestionRouter := adminRouter.Group("question")
			{
				adminQuestionRouter.GET(":id")
				adminQuestionRouter.PUT(":id")
				adminQuestionRouter.DELETE(":id")
			}
			adminRouter.POST("create-question")
		}
	}

	return router, nil
}
