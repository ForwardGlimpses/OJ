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

// Authentication 中间件：处理身份认证和角色认证
// requiredLevel 是需要的用户级别，只有级别高于或等于该级别的用户才能访问
func Authentication(requiredLevel int) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求中的 token
		token, err := c.Cookie("token")
		if err != nil {
			ginx.ResError(c, errors.AuthFailed("Token is required"))
			c.Abort() // 如果没有 token，终止后续处理
			return
		}

		ctx := c.Request.Context()

		// 根据 token 获取用户 ID
		userId, userLevel, err := loginSvc.GetUserInfo(ctx, token)
		if err != nil {
			ginx.ResError(c, errors.AuthFailed("Invalid token"))
			return
		}

		// 判断用户是否有权限访问该接口
		if userLevel < requiredLevel {
			ginx.ResError(c, errors.AuthFailed("User does not have the required permissions"))
			return
		}

		// 如果认证通过，设置用户 ID 到上下文中
		//ctx := c.Request.Context()
		ctx = context.WithValue(ctx, "user_id", userId)
		c.Request = c.Request.WithContext(ctx)
	}
}
