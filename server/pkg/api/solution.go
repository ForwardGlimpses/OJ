package api

import (
	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/ginx"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
	"github.com/gin-gonic/gin"
)

type SolutionAPI struct{}

//var solutionSvc service.SolutionServiceInterface = &service.SolutionService{}

// Get 获取解决方案信息
func (s *SolutionAPI) Get(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBindUri(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	item, err := solutionSvc.Get(id.ID)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// Create 创建解决方案，传递给 service 层时使用 Item 类型
func (s *SolutionAPI) Create(c *gin.Context) {
	var item schema.SolutionItem
	if err := c.ShouldBindJSON(&item); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	if _, err := solutionSvc.Create(&item); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "创建成功")
}

// Update 更新解决方案
func (s *SolutionAPI) Update(c *gin.Context) {
	var item schema.SolutionItem
	var id schema.ID
	if err := c.ShouldBindJSON(&item); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}
	if err := c.ShouldBindUri(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	if err := solutionSvc.Update(id.ID, &item); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "更新成功")
}

// Delete 删除解决方案
func (s *SolutionAPI) Delete(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBindUri(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	if err := solutionSvc.Delete(id.ID); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "删除成功")
}
