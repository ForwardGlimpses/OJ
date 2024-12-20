package api

import (
	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/ginx"
	"github.com/ForwardGlimpses/OJ/server/pkg/logs"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"

	//"github.com/ForwardGlimpses/OJ/server/pkg/service"
	"github.com/gin-gonic/gin"
)

type ContestProblemAPI struct{}


// Get 获取指定ID的比赛问题信息
func (a *ContestProblemAPI) Get(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBindUri(&id); err != nil {
		logs.Error("Failed to bind URI:", err)
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	item, err := contestProblemSvc.Get(id.ID)
	if err != nil {
		logs.Error("Failed to get contest problem with ID:", id.ID, "Error:", err)
		ginx.ResError(c, err)
		return
	}
	logs.Info("Successfully retrieved contest problem with ID:", id.ID)
	ginx.ResSuccess(c, item)
}

// Create 创建新的比赛问题
func (a *ContestProblemAPI) Create(c *gin.Context) {
	var item *schema.ContestProblemItem
	if err := c.ShouldBindJSON(item); err != nil {
		logs.Error("Failed to bind JSON:", err)
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	if _, err := contestProblemSvc.Create(item); err != nil {
		logs.Error("Failed to create contest problem:", err)
		ginx.ResError(c, err)
		return
	}
	logs.Info("Successfully created contest problem")
	ginx.ResSuccess(c, "创建成功")
}

// Update 更新指定ID的比赛问题信息
func (a *ContestProblemAPI) Update(c *gin.Context) {
	var item schema.ContestProblemItem
	if err := c.ShouldBindJSON(item); err != nil {
		logs.Error("Failed to bind JSON:", err)
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	if err := contestProblemSvc.Update(item.ID, &item); err != nil {
		logs.Error("Failed to update contest problem with ID:", item.ID, "Error:", err)
		ginx.ResError(c, err)
		return
	}
	logs.Info("Successfully updated contest problem with ID:", item.ID)
	ginx.ResSuccess(c, "更新成功")
}

// Delete 删除指定ID的比赛问题
func (a *ContestProblemAPI) Delete(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBindUri(&id); err != nil {
		logs.Error("Failed to bind URI:", err)
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	if err := contestProblemSvc.Delete(id.ID); err != nil {
		logs.Error("Failed to delete contest problem with ID:", id.ID, "Error:", err)
		ginx.ResError(c, err)
		return
	}
	logs.Info("Successfully deleted contest problem with ID:", id.ID)
	ginx.ResSuccess(c, "删除成功")
}
