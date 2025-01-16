package api

import (
	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/ginx"
	"github.com/ForwardGlimpses/OJ/server/pkg/service"
	"github.com/gin-gonic/gin"
)

type ContestSolutionAPI struct{}

// GetContestSolutions 获取比赛的所有解决方案
func (a *ContestSolutionAPI) GetContestSolutions(c *gin.Context) {
	var id struct {
		ContestID int `uri:"contest_id" binding:"required"`
	}
	if err := c.ShouldBindUri(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到比赛ID"))
		return
	}

	solutions, err := service.ContestSolutionSvc.GetContestSolutions(id.ContestID)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, solutions)
}

// GetContestRanking 获取比赛的实时排名
func (a *ContestSolutionAPI) GetContestRanking(c *gin.Context) {
	var id struct {
		ContestID int `uri:"contest_id" binding:"required"`
	}
	if err := c.ShouldBindUri(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到比赛ID"))
		return
	}

	ranking, err := service.ContestSolutionSvc.GetContestRanking(id.ContestID)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, ranking)
}
