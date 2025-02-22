package api

import (
	"fmt"
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

	ctx := c.Request.Context()

	submission := &schema.ContestSolutionItem{
		ContestID:   req.ContestID,
		ProblemID:   req.ProblemID,
		UserID:      req.UserID,
		SubmitTime:  time.Now(),
		IsAccepted:  false, // 初始状态为未接受
		PenaltyTime: 0,
	}

	id, err := service.ContestSolutionSvc.Create(ctx, submission)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResSuccess(c, id)
}

func (a *ContestSolutionAPI) Submit(c *gin.Context) {

	var input schema.Submit
	var id schema.ID
	// 绑定请求体数据到 input 结构体
	if err := c.ShouldBindJSON(&input); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	if err := c.ShouldBindUri(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}
	// 日志记录提交请求
	logs.Infof("用户 %d 提交了题目 %d 的代码", input.UserID, input.ID)

	ctx := c.Request.Context()

	// 调用 ProblemService 中的 Submit 方法，处理代码提交
	submissionID, err := service.ContestSolutionSvc.Submit(ctx, id.ID, &input)
	if err != nil {
		// 如果提交失败，记录并返回错误信息
		ginx.ResError(c, err)
		return
	}

	// 提交成功，返回提交 ID
	ginx.ResSuccess(c, submissionID)

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

	ctx := c.Request.Context()

	if err := contestSolutionSvc.Update(ctx, id.ID, &item); err != nil {
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

	ctx := c.Request.Context()

	if err := contestSolutionSvc.Delete(ctx, id.ID); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

// GetContestSolutions 获取比赛的所有解决方案
func (a *ContestSolutionAPI) Query(c *gin.Context) {
	var params schema.ContestSolutionParams
	if err := c.ShouldBindQuery(&params); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}
	ctx := c.Request.Context()

	items, total, err := contestSolutionSvc.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResSuccess(c, schema.QueryResult[schema.ContestSolutionItems]{
		Items:      items,
		TotalCount: total,
		Page:       params.Page,
		PageSize:   params.PageSize,
	})
}

// GetContestRanking 获取比赛的实时排名
func (a *ContestSolutionAPI) GetContestRanking(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBindUri(&id); err != nil {
		fmt.Println("contestsolution err: ", err)
		ginx.ResError(c, errors.InvalidInput("未找到比赛ID"))
		return
	}
	var params schema.ContestSolutionParams
	if err := c.ShouldBindQuery(&params); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	ctx := c.Request.Context()

	ranking, err := service.ContestSolutionSvc.GetContestRanking(ctx, id.ID, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, ranking)
}
