package api

import "github.com/ForwardGlimpses/OJ/server/pkg/service"

var (
	problemSvc service.ProblemServiceInterface
)

func init() {
	problemSvc = service.ProblemServiceInstance
}
