package initialize

import (
	"MSC2021/src/api"
	"MSC2021/src/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() (*gin.Engine, error) {
	router := gin.Default()

	apiRouter := router.Group("api")
	{
		accountRouter := apiRouter.Group("account")
		{
			accountRouter.POST("login", api.LoginHandler)
			accountRouter.POST("register", api.RegisterHandler)
			accountRouter.Use(middleware.LoginRequired())
			accountRouter.GET("me", api.GetProfileHandler)      // get profile
			accountRouter.PUT("me", api.ChangeProfileHandler)   // change profile
			accountRouter.POST("me", api.ChangePasswordHandler) // change password
			accountRouter.GET("logout", api.LogoutHandler)
		}
		apiRouter.GET("questions", api.GetQuestionListHandler)
		questionRouter := apiRouter.Group("question")
		{
			questionRouter.GET(":id", api.GetQuestionDetailHandler)
			questionRouter.Use(middleware.LoginRequired())
			questionRouter.POST(":id", api.AnswerQuestionHandler)
		}
		submissionRouter := apiRouter.Group("submission")
		{
			submissionRouter.GET(":quesId", api.GetSubmissionOfQuestionHandler) // get question [id] submissions of himself.
		}
		adminRouter := apiRouter.Group("admin")
		adminRouter.Use(middleware.LoginRequired())
		adminRouter.Use(middleware.AdminRequired())
		{
			adminRouter.GET("users", api.AdminGetUserListHandler)
			adminUserRouter := adminRouter.Group("user")
			{
				adminUserRouter.GET(":id", api.AdminGetUserProfileHandler)
				adminUserRouter.POST(":id", api.AdminChangeUserPasswordHandler)
				adminUserRouter.PUT(":id", api.AdminChangeUserProfileHandler)
				adminUserRouter.DELETE(":id", api.AdminDeleteUserHandler)
			}
			adminRouter.POST("create-user", api.AdminCreateUserHandler)
			adminRouter.GET("questions", api.AdminGetQuestionListHandler)
			adminQuestionRouter := adminRouter.Group("question")
			{
				adminQuestionRouter.GET(":id", api.AdminGetQuestionDetailHandler)
				adminQuestionRouter.PUT(":id", api.AdminChangeQuestionDetailHandler)
				adminQuestionRouter.DELETE(":id", api.AdminDeleteQuestionHandler)
			}
			adminRouter.POST("create-question", api.AdminCreateQuestionHandler)
			adminSubmissionRouter := adminRouter.Group("submission")
			{
				adminSubmissionRouter.GET(":quesId", api.AdminGetSubmittedUserOfQuestionHandler)
				adminSubmissionRouter.GET(":quesId/:userId", api.AdminGetSubmissionOfQuestionAndUserHandler) // get submissions of question [id]
			}
		}
	}

	return router, nil
}
