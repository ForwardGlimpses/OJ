package api

import (
	"context"

	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/ginx"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"

	//"github.com/ForwardGlimpses/OJ/server/pkg/service"
	"github.com/gin-gonic/gin"
)

type SourceCodeAPI struct{}

//var sourceCodeSvc service.SourceCodeServiceInterface = &service.SourceCodeService{}

func (a *SourceCodeAPI) Get(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBindUri(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	ctx := context.Background()

	item, err := sourceCodeSvc.Get(ctx, id.ID)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

func (a *SourceCodeAPI) Create(c *gin.Context) {
	var item schema.SourceCodeItem
	if err := c.ShouldBindJSON(&item); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	ctx := context.Background()

	if _, err := sourceCodeSvc.Create(ctx, &item); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "创建成功")
}

func (a *SourceCodeAPI) Update(c *gin.Context) {
	var item schema.SourceCodeItem
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

	if err := sourceCodeSvc.Update(ctx, id.ID, &item); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "更新成功")
}

func (a *SourceCodeAPI) Delete(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBindUri(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	ctx := context.Background()

	if err := sourceCodeSvc.Delete(ctx, id.ID); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "删除成功")
}
