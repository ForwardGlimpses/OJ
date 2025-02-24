package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ForwardGlimpses/OJ/server/pkg/logs"
	"github.com/ForwardGlimpses/OJ/server/pkg/service"
	"github.com/gin-gonic/gin"
)

type LoginAPI struct{}

// Login 用户登录
func (a *LoginAPI) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logs.Info("error: ", err)
		fmt.Println("error: ", err)
		return
	}

	ctx := context.Background()

	token, err := service.LoginSvc.Login(ctx, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		logs.Info("error: ", err)
		fmt.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
	fmt.Println("token: ", token)
}
