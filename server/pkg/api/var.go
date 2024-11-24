package api

import "github.com/ForwardGlimpses/OJ/server/pkg/service"

var (
	problemSvc        service.ProblemServiceInterface
	contestSvc        service.ContestServiceInterface
	contestProblemSvc service.ContestProblemServiceInterface
	solutionSvc       service.SolutionServiceInterface
	sourceCodeSvc     service.SourceCodeServiceInterface
	usersSvc          service.UsersServiceInterface
)

func init() {
	problemSvc = service.ProblemServiceInstance
	contestSvc = service.ContestServiceInstance
	contestProblemSvc = service.ContestProblemServiceInstance
	solutionSvc = service.SolutionServiceInstance
	sourceCodeSvc = service.SourceCodeServiceInstance
	usersSvc = service.UsersServiceInstance
}
