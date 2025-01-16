package bootstrap

import (
	"net/http"

	"github.com/ForwardGlimpses/OJ/server/pkg/global"
	"github.com/ForwardGlimpses/OJ/server/pkg/route"
	"github.com/gin-gonic/gin"
)

func Run() error {
	err := global.Init()
	if err != nil {
		return err
	}
	g := gin.New()

	// 配置 CORS
	// g.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:8080"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }))

	gApi := g.Group("api")

	route.RegisterProblem(gApi)
	route.RegisterContestSolution(gApi)
	route.RegisterContestProblem(gApi)
	route.RegisterContest(gApi)
	route.RegisterSourceCode(gApi)
	route.RegisterSolution(gApi)
	route.RegisterUsers(gApi)
	route.RegisterLogin(gApi)
	srv := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: g,
	}

	err = srv.ListenAndServe()
	return err
}
