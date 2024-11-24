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
	route.RegisterProblem(g.Group("api"))
	route.RegisterContestProblem(g.Group("api"))
	route.RegisterContest(g.Group("api"))
	route.RegisterSourceCode(g.Group("api"))
	route.RegisterSolution(g.Group("api"))
	route.RegisterUsers(g.Group("api"))
	srv := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: g,
	}

	err = srv.ListenAndServe()
	return err
}
