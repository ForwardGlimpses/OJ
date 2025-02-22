package api

import (
	"github.com/ForwardGlimpses/OJ/server/pkg/middleware"
	"github.com/gin-gonic/gin"
)

//   8080/api/problem/xxx

func RegisterProblem(g *gin.RouterGroup) {
	gGroup := g.Group("problem")
	api := ProblemAPI{}
	gGroup.GET(":id", api.Get)
	gGroup.GET("", api.Query)
	gGroup.POST(":id", middleware.Authentication(1), api.Submit)
	gGroup.POST("", middleware.Authentication(2), api.Create)
	gGroup.PUT(":id", middleware.Authentication(2), api.Update)
	gGroup.DELETE(":id", middleware.Authentication(2), api.Delete)
}

func RegisterContestSolution(g *gin.RouterGroup) {
	gGroup := g.Group("contestSolution")
	api := ContestSolutionAPI{}
	gGroup.GET(":id/rank", middleware.Authentication(1), api.GetContestRanking)
	gGroup.GET("", api.Query)
	gGroup.POST(":id", middleware.Authentication(1), api.Submit)
	gGroup.POST("", middleware.Authentication(2), api.Create)
	gGroup.PUT(":id", middleware.Authentication(2), api.Update)
	gGroup.DELETE(":id", middleware.Authentication(2), api.Delete)
}

func RegisterContestProblem(g *gin.RouterGroup) {
	gGroup := g.Group("contestProblem")
	api := ContestProblemAPI{}
	gGroup.GET("", api.Query)
	gGroup.GET(":id", api.Get)
	gGroup.POST("", middleware.Authentication(2), api.Create)
	gGroup.PUT(":id", middleware.Authentication(2), api.Update)
	gGroup.DELETE(":id", middleware.Authentication(2), api.Delete)
}

func RegisterContest(g *gin.RouterGroup) {
	gGroup := g.Group("contest")
	api := ContestAPI{}
	gGroup.GET(":id", api.Get)
	gGroup.GET("", api.Query)
	gGroup.POST("", middleware.Authentication(1), api.Create)
	gGroup.PUT(":id", middleware.Authentication(2), api.Update)
	gGroup.DELETE(":id", middleware.Authentication(2), api.Delete)
}

func RegisterSolution(g *gin.RouterGroup) {
	gGroup := g.Group("solution")
	api := SolutionAPI{}
	gGroup.GET(":id", api.Get)
	gGroup.GET("", api.Query)
	gGroup.POST("", middleware.Authentication(2), api.Create)
	gGroup.PUT(":id", middleware.Authentication(2), api.Update)
	gGroup.DELETE(":id", middleware.Authentication(2), api.Delete)
}

func RegisterSourceCode(g *gin.RouterGroup) {
	gGroup := g.Group("sourceCode")
	api := SourceCodeAPI{}
	gGroup.GET(":id", api.Get)
	gGroup.GET("", api.Query)
	gGroup.POST("", middleware.Authentication(1), api.Create)
	gGroup.PUT(":id", middleware.Authentication(2), api.Update)
	gGroup.DELETE(":id", middleware.Authentication(2), api.Delete)
}

func RegisterUsers(g *gin.RouterGroup) {
	gGroup := g.Group("users")
	api := UsersAPI{}
	gGroup.GET(":id", api.Get)
	gGroup.POST("register", api.Register)
	gGroup.POST("", api.Create)
	gGroup.PUT(":id", middleware.Authentication(2), api.Update)
	gGroup.DELETE(":id", middleware.Authentication(2), api.Delete)
}

// RegisterLogin 注册认证相关路由
func RegisterLogin(g *gin.RouterGroup) {
	gGroup := g.Group("login")
	api := LoginAPI{}
	gGroup.POST("", api.Login)
}
