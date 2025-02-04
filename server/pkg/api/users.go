package api

import (
	"context"

	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/ginx"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
	"github.com/gin-gonic/gin"
)

// UsersAPI 用户 API
type UsersAPI struct{}

// Get 获取用户信息
func (a *UsersAPI) Get(c *gin.Context) {
	var id struct {
		ID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的用户ID"))
		return
	}

	ctx := context.Background()

	item, err := usersSvc.Get(ctx, id.ID)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// Create 创建用户
func (a *UsersAPI) Create(c *gin.Context) {
	var item schema.UsersItem
	if err := c.ShouldBindJSON(&item); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	ctx := context.Background()

	id, err := usersSvc.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, id)
}

// Create 创建用户
func (a *UsersAPI) Register(c *gin.Context) {
	var item schema.UsersItem

	if err := c.ShouldBindJSON(&item); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	ctx := context.Background()

	item.Level = 1

	id, err := usersSvc.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, id)
}

// Update 更新用户信息
func (a *UsersAPI) Update(c *gin.Context) {
	var item schema.UsersItem
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

	if err := usersSvc.Update(ctx, id.ID, &item); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "更新成功")
}

// Delete 删除用户
func (a *UsersAPI) Delete(c *gin.Context) {
	var id struct {
		ID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的用户ID"))
		return
	}

	ctx := context.Background()

	if err := usersSvc.Delete(ctx, id.ID); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "删除成功")
}
