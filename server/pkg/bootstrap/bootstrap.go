package bootstrap

import (
	"net/http"

	"github.com/ForwardGlimpses/OJ/server/pkg/route"
	"github.com/gin-gonic/gin"
)

func Run() error {
	g := gin.New()
	route.RegisterProblem(g.Group("api"))

	srv := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: g,
	}

	err := srv.ListenAndServe()
	return err
}