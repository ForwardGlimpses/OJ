package api

import (
	"context"

	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/ginx"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"

	//"github.com/ForwardGlimpses/OJ/server/pkg/service"
	"github.com/gin-gonic/gin"
)

type ContestAPI struct{}

func (a *ContestAPI) Get(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBindUri(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	ctx := context.Background()

	item, err := contestSvc.Get(ctx, id.ID)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

func (a *ContestAPI) Create(c *gin.Context) {
	var item schema.ContestItem
	if err := c.ShouldBindJSON(&item); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	ctx := context.Background()

	id, err := contestSvc.Create(ctx, &item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, id)
}

func (a *ContestAPI) Update(c *gin.Context) {
	var item schema.ContestItem
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

	if err := contestSvc.Update(ctx, id.ID, &item); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

func (a *ContestAPI) Delete(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBindUri(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	ctx := context.Background()

	if err := contestSvc.Delete(ctx, id.ID); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
