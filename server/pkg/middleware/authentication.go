package middleware

import (
	"context"

	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/ginx"
	"github.com/ForwardGlimpses/OJ/server/pkg/service"
	"github.com/gin-gonic/gin"
)

var loginSvc service.LoginServiceInterface

func init() {
	loginSvc = service.LoginSvc
}

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			return
		}

		userId, err := loginSvc.GetUserId(token)
		if err != nil {
			ginx.ResError(c, errors.AuthFailed(err.Error()))
			return
		}

		// TODO: 处理不同权限的 API

		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, "user_id", userId)
		c.Request = c.Request.WithContext(ctx)
	}
}
