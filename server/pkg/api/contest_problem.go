package api

import (
	"fmt"

	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/ginx"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"

	//"github.com/ForwardGlimpses/OJ/server/pkg/service"
	"github.com/gin-gonic/gin"
)

type ContestProblemAPI struct{}

func (a *ContestProblemAPI) Query(c *gin.Context) {
	var params schema.ContestProblemParams
	if err := c.ShouldBindQuery(&params); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	ctx := c.Request.Context()

	items, total, err := contestProblemSvc.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResSuccess(c, schema.QueryResult[schema.ContestProblemItems]{
		Items:      items,
		TotalCount: total,
		Page:       params.Page,
		PageSize:   params.PageSize,
	})
}

// Get 获取指定ID的比赛问题信息
func (a *ContestProblemAPI) Get(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBindUri(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}
	ctx := c.Request.Context()
	item, err := contestProblemSvc.Get(ctx, id.ID)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// Create 创建新的比赛问题
func (a *ContestProblemAPI) Create(c *gin.Context) {
	var item schema.ContestProblemItem

	if err := c.ShouldBindJSON(&item); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		fmt.Println("item err ", err)
		return
	}

	ctx := c.Request.Context()

	id, err := contestProblemSvc.Create(ctx, &item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, id)
}

// Update 更新指定ID的比赛问题信息
func (a *ContestProblemAPI) Update(c *gin.Context) {
	var item schema.ContestProblemItem
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

	if err := contestProblemSvc.Update(ctx, id.ID, &item); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

// Delete 删除指定ID的比赛问题
func (a *ContestProblemAPI) Delete(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBindUri(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	ctx := c.Request.Context()

	if err := contestProblemSvc.Delete(ctx, id.ID); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
