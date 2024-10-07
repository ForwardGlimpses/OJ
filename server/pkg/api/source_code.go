package api

import (
	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/ginx"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
	"github.com/ForwardGlimpses/OJ/server/pkg/service"
	"github.com/gin-gonic/gin"
)

type SourceCodeAPI struct{}

var sourceCodeSvc service.SourceCodeServiceInterface = &service.SourceCodeService{}

// Get 获取指定ID的源代码信息
func (a *SourceCodeAPI) Get(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBind(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	item, err := sourceCodeSvc.Get(id.ID)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// Create 创建新的源代码
func (a *SourceCodeAPI) Create(c *gin.Context) {
	var item schema.Source_codeItem
	if err := c.ShouldBindJSON(&item); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	dbItem := schema.Source_codeDBItem{
		Solution_ID: item.Solution_ID,
		Source:      item.Source,
	}

	if err := sourceCodeSvc.Create(&dbItem); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "创建成功")
}

// Update 更新指定ID的源代码信息
func (a *SourceCodeAPI) Update(c *gin.Context) {
	var item schema.Source_codeItem
	if err := c.ShouldBindJSON(&item); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	dbItem := schema.Source_codeDBItem{
		Solution_ID: item.Solution_ID,
		Source:      item.Source,
	}

	if err := sourceCodeSvc.Update(item.Solution_ID, &dbItem); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "更新成功")
}

// Delete 删除指定ID的源代码
func (a *SourceCodeAPI) Delete(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBind(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	if err := sourceCodeSvc.Delete(id.ID); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "删除成功")
}
