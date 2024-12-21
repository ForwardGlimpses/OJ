package api

import (
	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/ginx"
	"github.com/ForwardGlimpses/OJ/server/pkg/logs"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
	"github.com/gin-gonic/gin"
)

type SolutionAPI struct{}

//var solutionSvc service.SolutionServiceInterface = &service.SolutionService{}

// Get 获取解决方案信息
func (s *SolutionAPI) Get(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBindUri(&id); err != nil {
		logs.Error("Failed to bind URI:", err)
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	item, err := solutionSvc.Get(id.ID)
	if err != nil {
		logs.Error("Failed to get solution with ID:", id.ID, "Error:", err)
		ginx.ResError(c, err)
		return
	}
	logs.Info("Successfully retrieved solution with ID:", id.ID)
	ginx.ResSuccess(c, item)
}

// Create 创建解决方案，传递给 service 层时使用 Item 类型
func (s *SolutionAPI) Create(c *gin.Context) {
	var item schema.SolutionItem
	if err := c.ShouldBindJSON(&item); err != nil {
		logs.Error("Failed to bind JSON:", err)
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	if _, err := solutionSvc.Create(&item); err != nil {
		logs.Error("Failed to create solution:", err)
		ginx.ResError(c, err)
		return
	}
	logs.Info("Successfully created solution")
	ginx.ResSuccess(c, "创建成功")
}

// Update 更新解决方案
func (s *SolutionAPI) Update(c *gin.Context) {
	var item schema.SolutionItem
	if err := c.ShouldBindJSON(&item); err != nil {
		logs.Error("Failed to bind JSON:", err)
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	if err := solutionSvc.Update(item.ID, &item); err != nil {
		logs.Error("Failed to update solution with ID:", item.ID, "Error:", err)
		ginx.ResError(c, err)
		return
	}
	logs.Info("Successfully updated solution with ID:", item.ID)
	ginx.ResSuccess(c, "更新成功")
}

// Delete 删除解决方案
func (s *SolutionAPI) Delete(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBindUri(&id); err != nil {
		logs.Error("Failed to bind URI:", err)
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	if err := solutionSvc.Delete(id.ID); err != nil {
		logs.Error("Failed to delete solution with ID:", id.ID, "Error:", err)
		ginx.ResError(c, err)
		return
	}
	logs.Info("Successfully deleted solution with ID:", id.ID)
	ginx.ResSuccess(c, "删除成功")
}
