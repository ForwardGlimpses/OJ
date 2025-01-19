package api

import (
	"time"

	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/ginx"
	"github.com/ForwardGlimpses/OJ/server/pkg/logs"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
	"github.com/ForwardGlimpses/OJ/server/pkg/service"
	"github.com/gin-gonic/gin"
)

type ContestSolutionAPI struct{}

// SubmitSolution 提交解决方案
func (a *ContestSolutionAPI) Create(c *gin.Context) {
	var req struct {
		ContestID int    `json:"contest_id"`
		ProblemID int    `json:"problem_id"`
		UserID    int    `json:"user_id"`
		Code      string `json:"code"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	submission := &schema.ContestSolutionItem{
		ContestID:   req.ContestID,
		ProblemID:   req.ProblemID,
		UserID:      req.UserID,
		SubmitTime:  time.Now(),
		IsAccepted:  false, // 初始状态为未接受
		PenaltyTime: 0,
	}

	id, err := service.ContestSolutionSvc.Create(submission)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResSuccess(c, id)
}

func (a *ContestSolutionAPI) Submit(c *gin.Context) {

	var input schema.Submit
	// 绑定请求体数据到 input 结构体
	if err := c.ShouldBindJSON(&input); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	// 日志记录提交请求
	logs.Infof("用户 %d 提交了比赛题目 %d 的代码", input.UserID, input.ID)

	// 调用 ProblemService 中的 Submit 方法，处理代码提交
	submissionID, err := service.ContestSolutionSvc.Submit(input.ID, input.UserID, input.InputCode)
	if err != nil {
		// 如果提交失败，记录并返回错误信息
		ginx.ResError(c, err)
		return
	}

	// 提交成功，返回提交 ID
	ginx.ResSuccess(c, gin.H{
		"submission_id": submissionID,
		"message":       "提交成功",
	})
}

// UpdateSolution 更新比赛解决方案信息
func (a *ContestSolutionAPI) Update(c *gin.Context) {
	var item schema.ContestSolutionItem
	var id schema.ID
	if err := c.ShouldBindJSON(&item); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}
	if err := c.ShouldBindUri(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	if err := contestSolutionSvc.Update(id.ID, &item); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

// DeleteSolution 删除比赛解决方案
func (a *ContestSolutionAPI) Delete(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBindUri(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	if err := contestSolutionSvc.Delete(id.ID); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

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
