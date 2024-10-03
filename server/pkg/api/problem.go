package api

import (
	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/ginx"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
	"github.com/gin-gonic/gin"
)

type ProblemAPI struct{}

func (a *ProblemAPI) Get(c *gin.Context) {
	var id schema.ID
	err := c.ShouldBind(&id)
	if err != nil {
		ginx.ResError(c, errors.InvalidInput("not found id"))
		return
	}

	item, err := problemSvc.Get(id.ID)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}
