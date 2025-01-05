package api

import (
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

	item, err := usersSvc.Get(id.ID)
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

	id, err := usersSvc.Create(item)
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

	item.Level = 1

	id, err := usersSvc.Create(item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, id)
}

// Update 更新用户信息
func (a *UsersAPI) Update(c *gin.Context) {
	var item schema.UsersItem
	if err := c.ShouldBindJSON(&item); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	if err := usersSvc.Update(item.ID, &item); err != nil {
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

	if err := usersSvc.Delete(id.ID); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "删除成功")
}
