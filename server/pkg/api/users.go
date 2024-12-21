package api

import (
	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/ginx"
	"github.com/ForwardGlimpses/OJ/server/pkg/logs"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"

	//"github.com/ForwardGlimpses/OJ/server/pkg/service"
	"github.com/gin-gonic/gin"
)

type UsersAPI struct{}

//var usersSvc service.UsersServiceInterface = &service.UsersService{}

func (a *UsersAPI) Get(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBindUri(&id); err != nil {
		logs.Error("Failed to bind URI:", err)
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	item, err := usersSvc.Get(id.ID)
	if err != nil {
		logs.Error("Failed to get users with ID:", id.ID, "Error:", err)
		ginx.ResError(c, err)
		return
	}
	logs.Info("Successfully retrieved users with ID:", id.ID)
	ginx.ResSuccess(c, item)
}

func (a *UsersAPI) Create(c *gin.Context) {
	var item schema.UsersItem
	if err := c.ShouldBindJSON(&item); err != nil {
		logs.Error("Failed to bind JSON:", err)
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	if _, err := usersSvc.Create(&item); err != nil {
		logs.Error("Failed to create users:", err)
		ginx.ResError(c, err)
		return
	}
	logs.Info("Successfully created users")
	ginx.ResSuccess(c, "创建成功")
}

func (a *UsersAPI) Update(c *gin.Context) {
	var item schema.UsersItem
	if err := c.ShouldBindJSON(&item); err != nil {
		logs.Error("Failed to bind JSON:", err)
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	if err := usersSvc.Update(item.ID, &item); err != nil {
		logs.Error("Failed to update users with ID:", item.ID, "Error:", err)
		ginx.ResError(c, err)
		return
	}
	logs.Info("Successfully updated users with ID:", item.ID)
	ginx.ResSuccess(c, "更新成功")
}

func (a *UsersAPI) Delete(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBindUri(&id); err != nil {
		logs.Error("Failed to bind URI:", err)
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	if err := usersSvc.Delete(id.ID); err != nil {
		logs.Error("Failed to delete users with ID:", id.ID, "Error:", err)
		ginx.ResError(c, err)
		return
	}
	logs.Info("Successfully deleted users with ID:", id.ID)
	ginx.ResSuccess(c, "删除成功")
}
