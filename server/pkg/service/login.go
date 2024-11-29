package service

import (
	"sync"
	"time"

	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/global"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtSecret = []byte("your-secure-secret-key") // JWT 密钥，可以存储在配置文件中

type LoginServiceInterface interface {
	Login(email, password string) (token string, err error)
	Logout(token string) (err error)
	GetUserId(token string) (userId int, err error)
	GetUserRoles(userId int) ([]string, error)
}

var LoginSvc LoginServiceInterface = &LoginService{}

// LoginService 结构体，管理登录逻辑
type LoginService struct {
	tokenMap sync.Map
}

type TokenInfo struct {
	UserId     int
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
func (a *LoginService) Login(email, password string) (string, error) {
	user, err := UserSvc.GetWithEmail(email) // 获取用户信息
	if err != nil {
		return "", err
	}

	// 验证密码
	if err := checkPassword(user.Password, password); err != nil {
		return "", errors.AuthFailed("password error")
	}

	// 生成 JWT token
	token, err := generateJWT(user.ID)
	if err != nil {
		return "", errors.InternalServer("failed to generate token")
	}

	// 存储 token 信息
	a.tokenMap.Store(token, TokenInfo{
		UserId:     user.ID,
		Expiration: time.Now().Add(time.Hour), // Token 有效期设置为 1 小时
	})
	return token, nil
}

// 用户登出，删除 Token 信息
func (a *LoginService) Logout(token string) error {
	_, ok := a.tokenMap.LoadAndDelete(token)
	if !ok {
		return errors.InvalidInput("login expired or invalid token")
	}
	return nil
}

// 获取用户 ID，验证 Token 是否有效
func (a *LoginService) GetUserId(token string) (int, error) {
	// 验证 JWT Token 是否有效并提取用户 ID
	userID, err := validateJWT(token)
	if err != nil {
		return 0, errors.AuthFailed("invalid or expired token")
	}

	// 检查 token 是否仍然有效
	value, ok := a.tokenMap.Load(token)
	if !ok {
		return 0, errors.AuthFailed("token not found")
	}

	tokenInfo := value.(TokenInfo)
	now := time.Now()
	if tokenInfo.Expiration.Before(now) {
		return 0, errors.AuthFailed("token expired")
	}

	// 延长 token 的有效期
	tokenInfo.Expiration = now.Add(time.Hour)
	a.tokenMap.Store(token, tokenInfo)

	return userID, nil
}

// 获取用户角色
func (a *LoginService) GetUserRoles(userId int) ([]string, error) {
	var roles []string

	// 使用 gorm 查询用户角色
	err := global.DB.Model(&schema.UsersDBItem{}).Where("id = ?", userId).Pluck("role", &roles).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.InternalServer("User not found or has no roles")
		}
		return nil, errors.InternalServer("Failed to retrieve user roles: " + err.Error())
	}

	return roles, nil
}
