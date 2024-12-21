package api

import (
	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/ginx"
	"github.com/ForwardGlimpses/OJ/server/pkg/logs"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"

	//"github.com/ForwardGlimpses/OJ/server/pkg/service"
	"github.com/gin-gonic/gin"
)

type ContestAPI struct{}


func (a *ContestAPI) Get(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBindUri(&id); err != nil {
		logs.Error("Failed to bind URI:", err)
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	item, err := contestSvc.Get(id.ID)
	if err != nil {
		logs.Error("Failed to get contest with ID:", id.ID, "Error:", err)
		ginx.ResError(c, err)
		return
	}
	logs.Info("Successfully retrieved contest with ID:", id.ID)
	ginx.ResSuccess(c, item)
}

func (a *ContestAPI) Create(c *gin.Context) {
	var item schema.ContestItem
	if err := c.ShouldBindJSON(&item); err != nil {
		logs.Error("Failed to bind JSON:", err)
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	if _, err := contestSvc.Create(&item); err != nil {
		logs.Error("Failed to create contest:", err)
		ginx.ResError(c, err)
		return
	}
	logs.Info("Successfully created contest")
	ginx.ResSuccess(c, "创建成功")
}

func (a *ContestAPI) Update(c *gin.Context) {
	var item schema.ContestItem
	if err := c.ShouldBindJSON(&item); err != nil {
		logs.Error("Failed to bind JSON:", err)
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	if err := contestSvc.Update(item.ID, &item); err != nil {
		logs.Error("Failed to update contest with ID:", item.ID, "Error:", err)
		ginx.ResError(c, err)
		return
	}
	logs.Info("Successfully updated contest with ID:", item.ID)
	ginx.ResSuccess(c, "更新成功")
}

func (a *ContestAPI) Delete(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBindUri(&id); err != nil {
		logs.Error("Failed to bind URI:", err)
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	if err := contestSvc.Delete(id.ID); err != nil {
		logs.Error("Failed to delete contest with ID:", id.ID, "Error:", err)
		ginx.ResError(c, err)
		return
	}
	logs.Info("Successfully deleted contest with ID:", id.ID)
	ginx.ResSuccess(c, "删除成功")
}
