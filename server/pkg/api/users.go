package api

import (
	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/ginx"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
	"github.com/ForwardGlimpses/OJ/server/pkg/service"
	"github.com/gin-gonic/gin"
)

type UsersAPI struct{}

var usersSvc service.UsersServiceInterface = &service.UsersService{}

// Get 获取指定ID的用户信息
func (a *UsersAPI) Get(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBind(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	user, err := usersSvc.Get(id.ID)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, user)
}

// Create 创建新用户
func (a *UsersAPI) Create(c *gin.Context) {
	var user schema.UsersItem
	if err := c.ShouldBindJSON(&user); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	dbUser := schema.UsersDBItem{
		Email:       user.Email,
		Sumbit:      user.Sumbit,
		Solve:       user.Solve,
		Password:    user.Password,
		School:      user.School,
		Access_time: user.Access_time,
		Enroll_time: user.Enroll_time,
	}

	if err := usersSvc.Create(&dbUser); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "创建成功")
}

// Update 更新指定ID的用户信息
func (a *UsersAPI) Update(c *gin.Context) {
	var user schema.UsersItem
	if err := c.ShouldBindJSON(&user); err != nil {
		ginx.ResError(c, errors.InvalidInput("无效的输入数据"))
		return
	}

	dbUser := schema.UsersDBItem{
		User_ID:     user.User_ID,
		Email:       user.Email,
		Sumbit:      user.Sumbit,
		Solve:       user.Solve,
		Password:    user.Password,
		School:      user.School,
		Access_time: user.Access_time,
		Enroll_time: user.Enroll_time,
	}

	if err := usersSvc.Update(user.User_ID, &dbUser); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "更新成功")
}

// Delete 删除指定ID的用户
func (a *UsersAPI) Delete(c *gin.Context) {
	var id schema.ID
	if err := c.ShouldBind(&id); err != nil {
		ginx.ResError(c, errors.InvalidInput("未找到ID"))
		return
	}

	if err := usersSvc.Delete(id.ID); err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, "删除成功")
}
