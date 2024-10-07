package api

import (
	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/ginx"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
	"github.com/ForwardGlimpses/OJ/server/pkg/service"
	"github.com/gin-gonic/gin"
)

type ContestAPI struct{}

var contestSvc service.ContestServiceInterface = &service.ContestService{}

// Get 获取指定ID的比赛信息
func (a *ContestAPI) Get(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBind(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	item, err := contestSvc.Get(id.ID)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// Create 创建新的比赛
func (a *ContestAPI) Create(c *gin.Context) {
	var item schema.ContestItem
	if err := c.ShouldBindJSON(&item); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	dbItem := schema.ContestDBItem{
		Title:        item.Title,
		Start_time:   item.Start_time,
		End_time:     item.End_time,
		Password:     item.Password,
		Administrator: item.Administrator,
		Description:  item.Description,
	}

	if err := contestSvc.Create(&dbItem); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "创建成功")
}

// Update 更新指定ID的比赛信息
func (a *ContestAPI) Update(c *gin.Context) {
	var item schema.ContestItem
	if err := c.ShouldBindJSON(&item); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	dbItem := schema.ContestDBItem{
		Contest_ID:   item.Contest_ID,
		Title:        item.Title,
		Start_time:   item.Start_time,
		End_time:     item.End_time,
		Password:     item.Password,
		Administrator: item.Administrator,
		Description:  item.Description,
	}

	if err := contestSvc.Update(item.Contest_ID, &dbItem); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "更新成功")
}

// Delete 删除指定ID的比赛
func (a *ContestAPI) Delete(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBind(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	if err := contestSvc.Delete(id.ID); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "删除成功")
}
