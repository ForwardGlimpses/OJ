package service

import (
	"context"
	"sync"
	"time"

	"github.com/ForwardGlimpses/OJ/server/pkg/config"
	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/logs"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your-secure-secret-key") // JWT 密钥，可以存储在配置文件中

type LoginServiceInterface interface {
	Login(ctx context.Context, email, password string) (token string, err error)
	Logout(ctx context.Context, token string) (err error)
	GetUserInfo(ctx context.Context, token string) (userId int, userLevel int, err error)
}

var LoginSvc LoginServiceInterface = &LoginService{}

// LoginService 结构体，管理登录逻辑
type LoginService struct {
	tokenMap sync.Map
}

type TokenInfo struct {
	UserId     int
	Level      int
	Expiration time.Time
}

// 验证密码
func checkPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// 生成 JWT Token
func generateJWT(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour).Unix(), // Token 有效期为1小时
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// 验证 JWT Token 并获取用户 ID
func validateJWT(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return 0, errors.AuthFailed("invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["user_id"] == nil {
		return 0, errors.AuthFailed("invalid token claims")
	}

	userID := int(claims["user_id"].(float64))
	return userID, nil
}

// 用户登录，生成 JWT Token
func (a *LoginService) Login(ctx context.Context, email, password string) (string, error) {

	var user *schema.UsersItem
	var err error
	if email == config.C.Root.Email && password == config.C.Root.Password {
		logs.Info("Root login")
		user = &schema.UsersItem{
			ID:       0,
			Email:    config.C.Root.Email,
			Password: config.C.Root.Password,
			Level:    3,
		}
	} else {
		logs.Info("User login")
		user, err = UserSvc.GetWithEmail(ctx, email) // 获取用户信息
		if err != nil {
			return "", err
		}
		// 验证密码
		if err := checkPassword(user.Password, password); err != nil {
			return "", errors.AuthFailed("password error")
		}
	}

	// 生成 JWT token
	token, err := generateJWT(user.ID)
	if err != nil {
		return "", errors.InternalServer("failed to generate token")
	}

	// 存储 token 信息
	a.tokenMap.Store(token, TokenInfo{
		UserId:     user.ID,
		Level:      user.Level,
		Expiration: time.Now().Add(time.Hour), // Token 有效期设置为 1 小时

	})
	return token, nil
}

// 用户登出，删除 Token 信息
func (a *LoginService) Logout(ctx context.Context, token string) error {
	_, ok := a.tokenMap.LoadAndDelete(token)
	if !ok {
		return errors.InvalidInput("login expired or invalid token")
	}
	return nil
}

// 获取用户 ID，验证 Token 是否有效
func (a *LoginService) GetUserInfo(ctx context.Context, token string) (int, int, error) {
	// 验证 JWT Token 是否有效并提取用户 ID
	_, err := validateJWT(token)
	if err != nil {
		return 0, 0, errors.AuthFailed("invalid or expired token")
	}

	// 检查 token 是否仍然有效
	value, ok := a.tokenMap.Load(token)
	if !ok {
		return 0, 0, errors.AuthFailed("token not found")
	}

	tokenInfo := value.(TokenInfo)
	now := time.Now()
	if tokenInfo.Expiration.Before(now) {
		return 0, 0, errors.AuthFailed("token expired")
	}

	// 延长 token 的有效期
	tokenInfo.Expiration = now.Add(time.Hour)
	a.tokenMap.Store(token, tokenInfo)

	return tokenInfo.UserId, tokenInfo.Level, nil
}
