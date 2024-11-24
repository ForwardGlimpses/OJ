package service

import (
	"sync"
	"time"

	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/google/uuid"
)

type LoginServiceInterface interface {
	Login(email, password string) (token string, err error)
	Logout(token string) (err error)
	GetUserId(token string) (userId int, err error)
}

var LoginSvc LoginServiceInterface = &LoginService{}

type LoginService struct {
	tokenMap sync.Map
}

type TokenInfo struct {
	UserId     int
	Expiration time.Time
}

func (a *LoginService) Login(email, password string) (string, error) {
	user, err := UserSvc.GetWithEmail(email)
	if err != nil {
		return "", err
	}

	// TODO: 使用密文存储 password
	if user.Password != password {
		return "", errors.AuthFailed("password error")
	}

	// TODO: 更好的生成 token
	token := uuid.New().String()
	a.tokenMap.Store(token, TokenInfo{
		UserId:     user.ID,
		Expiration: time.Now().Add(time.Hour), // 超时时间设置一个小时
	})
	return token, nil
}

func (a *LoginService) Logout(token string) error {
	_, ok := a.tokenMap.LoadAndDelete(token)
	if !ok {
		return errors.InvalidInput("login expired")
	}
	return nil
}

func (a *LoginService) GetUserId(token string) (int, error) {
	value, ok := a.tokenMap.Load(token)
	if !ok {
		return 0, errors.AuthFailed("login expired")
	}

	tokenInfo := value.(TokenInfo)
	now := time.Now()
	if tokenInfo.Expiration.Before(now) {
		return 0, errors.AuthFailed("login expired")
	}

	tokenInfo.Expiration = now.Add(time.Hour)
	a.tokenMap.Store(token, tokenInfo)
	return tokenInfo.UserId, nil
}
