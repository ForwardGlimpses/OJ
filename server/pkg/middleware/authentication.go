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
// requiredRoles 是需要的角色列表，只有角色匹配的用户才可以访问
func Authentication(requiredRoles []string) gin.HandlerFunc {
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

		// 获取用户的角色信息
		roles, err := loginSvc.GetUserRoles(userId)
		if err != nil {
			ginx.ResError(c, errors.AuthFailed("Unable to fetch user roles"))
			c.Abort() // 如果获取用户角色失败，终止后续处理
			return
		}

		// 判断用户是否有权限访问该接口
		if !isRoleAuthorized(roles, requiredRoles) {
			ginx.ResError(c, errors.AuthFailed("User does not have the required permissions"))
			c.Abort() // 如果没有权限，终止后续处理
			return
		}

		// 如果认证和权限检查通过，设置用户 ID 到上下文中
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, "user_id", userId)
		c.Request = c.Request.WithContext(ctx)

		// 继续处理请求
		c.Next()
	}
}

// isRoleAuthorized 检查用户角色是否有权限访问指定的 API
func isRoleAuthorized(userRoles []string, requiredRoles []string) bool {
	// 判断用户角色是否与需要的角色匹配
	for _, userRole := range userRoles {
		for _, requiredRole := range requiredRoles {
			if userRole == requiredRole {
				return true
			}
		}
	}
	return false
}
