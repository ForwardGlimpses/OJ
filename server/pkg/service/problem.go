package service

import "github.com/ForwardGlimpses/OJ/server/pkg/schema"

type ProblemServiceInterface interface {
	Get(id int) (schema.ProblemItem, error)
}

var ProblemServiceInstance ProblemServiceInterface = &ProblemService{}

type ProblemService struct{}

func (a *ProblemService) Get(id int) (schema.ProblemItem, error) {
	return schema.ProblemItem{}, nil
}
