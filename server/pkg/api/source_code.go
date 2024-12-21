package api

import (
	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/ginx"
	"github.com/ForwardGlimpses/OJ/server/pkg/logs"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"

	//"github.com/ForwardGlimpses/OJ/server/pkg/service"
	"github.com/gin-gonic/gin"
)

type SourceCodeAPI struct{}

//var sourceCodeSvc service.SourceCodeServiceInterface = &service.SourceCodeService{}

func (a *SourceCodeAPI) Get(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBindUri(&id); err != nil {
		logs.Error("Failed to bind URI:", err)
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	item, err := sourceCodeSvc.Get(id.ID)
	if err != nil {
		logs.Error("Failed to get source code with ID:", id.ID, "Error:", err)
		ginx.ResError(c, err)
		return
	}
	logs.Info("Successfully retrieved source code with ID:", id.ID)
	ginx.ResSuccess(c, item)
}

func (a *SourceCodeAPI) Create(c *gin.Context) {
	var item schema.SourceCodeItem
	if err := c.ShouldBindJSON(&item); err != nil {
		logs.Error("Failed to bind JSON:", err)
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	if _, err := sourceCodeSvc.Create(&item); err != nil {
		logs.Error("Failed to create source code:", err)
		ginx.ResError(c, err)
		return
	}
	logs.Info("Successfully created source code")
	ginx.ResSuccess(c, "创建成功")
}

func (a *SourceCodeAPI) Update(c *gin.Context) {
	var item schema.SourceCodeItem
	if err := c.ShouldBindJSON(&item); err != nil {
		logs.Error("Failed to bind JSON:", err)
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	if err := sourceCodeSvc.Update(item.ID, &item); err != nil {
		logs.Error("Failed to update source code with ID:", item.ID, "Error:", err)
		ginx.ResError(c, err)
		return
	}
	logs.Info("Successfully updated source code with ID:", item.ID)
	ginx.ResSuccess(c, "更新成功")
}

func (a *SourceCodeAPI) Delete(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBindUri(&id); err != nil {
		logs.Error("Failed to bind URI:", err)
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	if err := sourceCodeSvc.Delete(id.ID); err != nil {
		logs.Error("Failed to delete source code with ID:", id.ID, "Error:", err)
		ginx.ResError(c, err)
		return
	}
	logs.Info("Successfully deleted source code with ID:", id.ID)
	ginx.ResSuccess(c, "删除成功")
}
