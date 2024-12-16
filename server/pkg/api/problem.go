package api

import (
	"log"

	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/ginx"
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

	item, err := problemSvc.Get(id.ID)
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

	if _, err := problemSvc.Create(&item); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "创建成功")
}

// Update 更新题目
func (a *ProblemAPI) Update(c *gin.Context) {
	var item schema.ProblemItem
	if err := c.ShouldBindJSON(&item); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	if err := problemSvc.Update(item.ID, &item); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "更新成功")
}

// Delete 删除题目
func (a *ProblemAPI) Delete(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBindUri(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	if err := problemSvc.Delete(id.ID); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "删除成功")
}

func (a *ProblemAPI) Submit(c *gin.Context) {
	// 从请求中获取提交数据
	var input struct {
		ProblemID int    `json:"id"`     // 题目 ID
		UserID    int    `json:"userid"` // 用户 ID
		Code      string `json:"code"`   // 用户提交的代码
	}

	// 绑定请求体数据到 input 结构体
	if err := c.ShouldBindJSON(&input); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	// 日志记录提交请求
	log.Printf("用户 %d 提交了题目 %d 的代码", input.UserID, input.ProblemID)

	// 调用 ProblemService 中的 Submit 方法，处理代码提交
	submissionID, err := service.ProblemSvc.Submit(input.ProblemID, input.UserID, input.Code)
	if err != nil {
		// 如果提交失败，记录并返回错误信息
		log.Printf("提交失败，用户 %d 提交题目 %d 时发生错误: %s", input.UserID, input.ProblemID, err.Error())
		ginx.ResError(c, err)
		return
	}

	// 提交成功，返回提交 ID
	log.Printf("用户 %d 成功提交了题目 %d,提交 ID: %d", input.UserID, input.ProblemID, submissionID)
	ginx.ResSuccess(c, gin.H{
		"submission_id": submissionID,
		"message":       "提交成功",
	})
}
