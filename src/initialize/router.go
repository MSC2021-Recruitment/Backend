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
			// 登录（签发token
			accountRouter.POST("login", api.LoginHandler)
			// 注册（不签发token
			accountRouter.POST("register", api.RegisterHandler)
			accountRouter.Use(middleware.LoginRequired())
			// 获取自己的信息
			accountRouter.GET("me", api.GetProfileHandler)
			// 改自己的信息
			accountRouter.PUT("me", api.ChangeProfileHandler)
			// 更改密码
			accountRouter.POST("me", api.ChangePasswordHandler)
			// 退出登录（主动将jwt token expire
			accountRouter.GET("logout", api.LogoutHandler)
		}
		// 获取问题列表
		apiRouter.GET("questions", api.GetQuestionListHandler)
		questionRouter := apiRouter.Group("question")
		{
			// 获取问题详情
			questionRouter.GET(":id", api.GetQuestionDetailHandler)
			questionRouter.Use(middleware.LoginRequired())
			// 回答问题
			questionRouter.POST(":id", api.AnswerQuestionHandler)
		}
		submissionRouter := apiRouter.Group("submission")
		{
			// 获取对应题目下他自己的submission
			submissionRouter.GET(":quesId", api.GetSubmissionOfQuestionHandler)
		}
		apiRouter.GET("groups", api.GetGroupListHandler)
		groupRouter := apiRouter.Group("group")
		{
			// ~~获取组别信息~~ 有个屁的组别信息，有名字就够了
			// groupRouter.GET(":id")
			groupRouter.Use(middleware.LoginRequired())
			// 当前登录的user加入一个组
			groupRouter.POST(":id", api.JoinGroupHandler)
			// 当前登录user离开一个组
			groupRouter.DELETE(":id", api.LeaveGroupHandler)
		}
		apiRouter.Use(middleware.LoginRequired())
		apiRouter.POST("sign", api.UserSignInInterviewHandler) // 面试签到
		adminRouter := apiRouter.Group("admin")
		adminRouter.Use(middleware.AdminRequired())
		{
			// 获取用户列表
			adminRouter.GET("users", api.AdminGetUserListHandler)
			adminUserRouter := adminRouter.Group("user")
			{
				// 获取用户的详细信息（所有
				adminUserRouter.GET(":id", api.AdminGetUserProfileHandler)
				// 更改用户的密码（只能重置
				adminUserRouter.POST(":id", api.AdminChangeUserPasswordHandler)
				// 更改用户的信息
				adminUserRouter.PUT(":id", api.AdminChangeUserProfileHandler)
				// 删除一个用户
				adminUserRouter.DELETE(":id", api.AdminDeleteUserHandler)
			}
			// 获取所有question（和答了这道题的user数量
			adminRouter.GET("questions", api.AdminGetQuestionListHandler)
			adminQuestionRouter := adminRouter.Group("question")
			{
				// 获取题目详情
				adminQuestionRouter.GET(":id", api.AdminGetQuestionDetailHandler)
				// 更新题目描述
				adminQuestionRouter.PUT(":id", api.AdminChangeQuestionDetailHandler)
				// 删除题目
				adminQuestionRouter.DELETE(":id", api.AdminDeleteQuestionHandler)
			}
			// 创建一个新的question
			adminRouter.POST("create-question", api.AdminCreateQuestionHandler)
			adminSubmissionRouter := adminRouter.Group("submission")
			{
				// 获取在此question上所有交过答案的user
				adminSubmissionRouter.GET(":quesId", api.AdminGetSubmittedUserOfQuestionHandler)
				// 获取user在此question上的所有提交
				adminSubmissionRouter.GET(":quesId/:userId", api.AdminGetSubmissionOfQuestionAndUserHandler)
			}
			// 获取所有组的信息（包括每个组的报名人数
			adminRouter.GET("groups", api.AdminGetGroupListHandler)
			adminGroupRouter := adminRouter.Group("group")
			{
				// 获取组信息（包括该组报名的人
				adminGroupRouter.GET(":groupId", api.AdminGetGroupDetailHandler)
				// 更改组信息
				adminGroupRouter.PUT(":groupId", api.AdminChangeGroupDetailHandler)
				// 删除一个组
				adminGroupRouter.DELETE(":groupId", api.AdminDeleteGroupHandler)
			}
			adminRouter.POST("create-group", api.AdminCreateGroupHandler) // 创建一个新组
			adminInterviewRouter := adminRouter.Group("interview")
			{
				// 获取建议面试的user（面试的优先队列实现
				adminInterviewRouter.GET(":groupId", api.AdminGetSuggestInterviewUserHandler)
				// 建议面试
				adminInterviewRouter.POST(":groupId", api.AdminInterviewHandler)
				// 获取面试信息
				adminInterviewRouter.GET(":groupId/:userId", api.AdminGetInterviewContentHandler)
				// 新建/更新面试信息
				adminInterviewRouter.POST(":groupId/:userId", api.AdminUpdateInterviewContentHandler)
				// 删除面试信息
				adminInterviewRouter.DELETE(":groupId/:userId", api.AdminDeleteInterviewContentHandler)
			}
			// 获取当前面试详情
			adminRouter.GET("interview-status", api.GetInterviewStatusHandler)
			// 开始面试
			adminRouter.POST("sign", api.StartInterviewHandler)
			adminRouter.DELETE("sign", api.StopInterviewHandler)
		}
	}

	return router, nil
}
