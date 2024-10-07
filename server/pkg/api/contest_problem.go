package api

import (
	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/ginx"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
	"github.com/ForwardGlimpses/OJ/server/pkg/service"
	"github.com/gin-gonic/gin"
)

type ContestProblemAPI struct{}

var contestProblemSvc service.ContestProblemServiceInterface = &service.ContestProblemService{}

// Get 获取指定ID的比赛问题信息
func (a *ContestProblemAPI) Get(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBind(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	item, err := contestProblemSvc.Get(id.ID)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// Create 创建新的比赛问题
func (a *ContestProblemAPI) Create(c *gin.Context) {
	var item schema.Contest_ProblemItem
	if err := c.ShouldBindJSON(&item); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	dbItem := schema.Contest_ProblemDBItem{
		Problem_ID: item.Problem_ID,
		Contest_ID: item.Contest_ID,
		Title:      item.Title,
		Num:        item.Num,
		Accepted:   item.Accepted,
		Submit:     item.Submit,
	}

	if err := contestProblemSvc.Create(&dbItem); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "创建成功")
}

// Update 更新指定ID的比赛问题信息
func (a *ContestProblemAPI) Update(c *gin.Context) {
	var item schema.Contest_ProblemItem
	if err := c.ShouldBindJSON(&item); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	dbItem := schema.Contest_ProblemDBItem{
		Problem_ID: item.Problem_ID,
		Contest_ID: item.Contest_ID,
		Title:      item.Title,
		Num:        item.Num,
		Accepted:   item.Accepted,
		Submit:     item.Submit,
	}

	if err := contestProblemSvc.Update(item.Problem_ID, &dbItem); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "更新成功")
}

// Delete 删除指定ID的比赛问题
func (a *ContestProblemAPI) Delete(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBind(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	if err := contestProblemSvc.Delete(id.ID); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "删除成功")
}
