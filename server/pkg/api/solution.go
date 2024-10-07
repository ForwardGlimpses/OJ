package api

import (
	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/ginx"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
	"github.com/ForwardGlimpses/OJ/server/pkg/service"
	"github.com/gin-gonic/gin"
)

type SolutionAPI struct{}

var solutionSvc service.SolutionServiceInterface = &service.SolutionService{}

// Get 获取指定ID的解决方案信息
func (a *SolutionAPI) Get(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBind(&id); err != nil {
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

// Create 创建新的解决方案
func (a *SolutionAPI) Create(c *gin.Context) {
	var item schema.SolutionItem
	if err := c.ShouldBindJSON(&item); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	dbItem := schema.SolutionDBItem{
		Problem_ID:  item.Problem_ID,
		User_ID:     item.User_ID,
		Time:        item.Time,
		Memory:      item.Memory,
		In_date:     item.In_date,
		Language:    item.Language,
		Code_length: item.Code_length,
		Juage_time:  item.Juage_time,
		Juager:      item.Juager,
		Pass_rate:   item.Pass_rate,
	}

	if err := solutionSvc.Create(&dbItem); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "创建成功")
}

// Update 更新指定ID的解决方案信息
func (a *SolutionAPI) Update(c *gin.Context) {
	var item schema.SolutionItem
	if err := c.ShouldBindJSON(&item); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	dbItem := schema.SolutionDBItem{
		Solution_ID: item.Solution_ID,
		Problem_ID:  item.Problem_ID,
		User_ID:     item.User_ID,
		Time:        item.Time,
		Memory:      item.Memory,
		In_date:     item.In_date,
		Language:    item.Language,
		Code_length: item.Code_length,
		Juage_time:  item.Juage_time,
		Juager:      item.Juager,
		Pass_rate:   item.Pass_rate,
	}

	if err := solutionSvc.Update(item.Solution_ID, &dbItem); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "更新成功")
}

// Delete 删除指定ID的解决方案
func (a *SolutionAPI) Delete(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBind(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	if err := solutionSvc.Delete(id.ID); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "删除成功")
}
