package route

import (
	"github.com/ForwardGlimpses/OJ/server/pkg/api"
	"github.com/gin-gonic/gin"
)

func RegisterProblem(g *gin.RouterGroup) {
	gGroup := g.Group("problem")
	api := api.ProblemAPI{}
	gGroup.GET("", api.Get)
}
