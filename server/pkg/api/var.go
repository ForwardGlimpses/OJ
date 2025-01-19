package api

import "github.com/ForwardGlimpses/OJ/server/pkg/service"

var (
	problemSvc         service.ProblemServiceInterface
	contestSvc         service.ContestServiceInterface
	contestSolutionSvc service.ContestSolutionServiceInterface
	contestProblemSvc  service.ContestProblemServiceInterface
	solutionSvc        service.SolutionServiceInterface
	sourceCodeSvc      service.SourceCodeServiceInterface
	usersSvc           service.UsersServiceInterface
)

func init() {
	problemSvc = service.ProblemSvc
	contestSvc = service.ContestSvc
	contestSolutionSvc = service.ContestSolutionSvc
	contestProblemSvc = service.ContestProblemSvc
	solutionSvc = service.SolutionSvc
	sourceCodeSvc = service.SourceCodeSvc
	usersSvc = service.UserSvc
}
