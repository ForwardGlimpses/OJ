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

		// 根据 token 获取用户 ID
		userId, err := loginSvc.GetUserId(token)
		if err != nil {
			ginx.ResError(c, errors.AuthFailed("Invalid token"))
			c.Abort() // 如果获取用户 ID 失败，终止后续处理
			return
		}

		// 获取用户的 Level 信息
		userLevel, err := loginSvc.GetUserLevel(userId)
		if err != nil {
			ginx.ResError(c, errors.AuthFailed("Unable to fetch user level"))
			c.Abort() // 如果获取用户 level 失败，终止后续处理
			return
		}

		// 判断用户是否有权限访问该接口
		if userLevel < requiredLevel {
			ginx.ResError(c, errors.AuthFailed("User does not have the required permissions"))
			c.Abort() // 如果用户权限不足，终止后续处理
			return
		}

		// 如果认证通过，设置用户 ID 到上下文中
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, "user_id", userId)
		c.Request = c.Request.WithContext(ctx)

		// 继续处理请求
		c.Next()
	}
}
