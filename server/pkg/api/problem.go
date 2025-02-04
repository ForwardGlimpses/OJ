package api

import (
	"context"

	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/ginx"
	"github.com/ForwardGlimpses/OJ/server/pkg/logs"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
	"github.com/ForwardGlimpses/OJ/server/pkg/service"
	"github.com/gin-gonic/gin"
)

type ProblemAPI struct{}

// Get 获取题目信息
func (a *ProblemAPI) Get(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBindUri(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	ctx := context.Background()

	item, err := problemSvc.Get(ctx, id.ID)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// Create 创建题目，传递给 service 层时使用 Item 类型
func (a *ProblemAPI) Create(c *gin.Context) {
	var item schema.ProblemItem
	if err := c.ShouldBindJSON(&item); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	ctx := context.Background()

	id, err := problemSvc.Create(ctx, &item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, id)
}

// Update 更新题目
func (a *ProblemAPI) Update(c *gin.Context) {
	var item schema.ProblemItem
	var id schema.ID
	if err := c.ShouldBindJSON(&item); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}
	if err := c.ShouldBindUri(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	ctx := context.Background()

	if err := problemSvc.Update(ctx, id.ID, &item); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

// Delete 删除题目
func (a *ProblemAPI) Delete(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBindUri(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	ctx := context.Background()

	if err := problemSvc.Delete(ctx, id.ID); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

func (a *ProblemAPI) Submit(c *gin.Context) {

	var input schema.Submit
	// 绑定请求体数据到 input 结构体
	if err := c.ShouldBindJSON(&input); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	// 日志记录提交请求
	logs.Infof("用户 %d 提交了题目 %d 的代码", input.UserID, input.ID)

	// 调用 ProblemService 中的 Submit 方法，处理代码提交

	// ctx := context.Background()
	// ctx = context.WithValue(ctx, "problemID", input.ID)
	// cctx := context.WithValue(ctx, "userID", input.UserID)
	// ccctx := cctx.Value("problemID")

	ctx := context.Background()

	submissionID, err := service.ProblemSvc.Submit(ctx, input.ID, input.UserID, input.InputCode)
	if err != nil {
		// 如果提交失败，记录并返回错误信息
		ginx.ResError(c, err)
		return
	}

	// 提交成功，返回提交 ID
	ginx.ResSuccess(c, submissionID)
}
