package route

import (
	"github.com/ForwardGlimpses/OJ/server/pkg/api"
	"github.com/gin-gonic/gin"
)

//   8080/api/problem/xxx

func RegisterProblem(g *gin.RouterGroup) {
	gGroup := g.Group("problem")
	api := api.ProblemAPI{}
	gGroup.GET(":id", api.Get)
	gGroup.Use()
	gGroup.POST(":id", api.Submit)
	gGroup.POST("", api.Create)
	gGroup.PUT(":id", api.Update)
	gGroup.DELETE(":id", api.Delete)
}

func RegisterContestProblem(g *gin.RouterGroup) {
	gGroup := g.Group("contestProblem")
	api := api.ContestProblemAPI{}
	gGroup.GET(":id", api.Get)
	gGroup.Use()
	gGroup.POST("", api.Create)
	gGroup.PUT("", api.Update)
	gGroup.DELETE(":id", api.Delete)
}

func RegisterContest(g *gin.RouterGroup) {
	gGroup := g.Group("contest")
	api := api.ContestAPI{}
	gGroup.GET(":id", api.Get)
	gGroup.Use()
	gGroup.POST("", api.Create)
	gGroup.PUT("", api.Update)
	gGroup.DELETE(":id", api.Delete)
}

func RegisterSolution(g *gin.RouterGroup) {
	gGroup := g.Group("solution")
	api := api.SolutionAPI{}
	gGroup.GET(":id", api.Get)
	gGroup.Use()
	gGroup.POST("", api.Create)
	gGroup.PUT("", api.Update)
	gGroup.DELETE(":id", api.Delete)
}

func RegisterSourceCode(g *gin.RouterGroup) {
	gGroup := g.Group("sourceCode")
	api := api.SourceCodeAPI{}
	gGroup.GET(":id", api.Get)
	gGroup.Use()
	gGroup.POST("", api.Create)
	gGroup.PUT("", api.Update)
	gGroup.DELETE(":id", api.Delete)
}

func RegisterUsers(g *gin.RouterGroup) {
	gGroup := g.Group("users")
	api := api.UsersAPI{}
	gGroup.GET(":id", api.Get)
	gGroup.Use()
	gGroup.POST("", api.Create)
	gGroup.PUT("", api.Update)
	gGroup.DELETE(":id", api.Delete)
}
