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
	problemSvc = service.ProblemSvc
	contestSvc = service.ContestSvc
	contestProblemSvc = service.ContestProblemSvc
	solutionSvc = service.SolutionSvc
	sourceCodeSvc = service.SourceCodeSvc
	usersSvc = service.UserSvc
}
