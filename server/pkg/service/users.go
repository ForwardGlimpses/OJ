package service

import (
	"errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
)

type UsersServiceInterface interface {
	Get(id int) (schema.UsersItem, error)
	Create(item *schema.UsersDBItem) error
	Update(id int, item *schema.UsersDBItem) error
	Delete(id int) error
}

var UsersServiceInstance UsersServiceInterface = &UsersService{}

type UsersService struct{}

// 模拟数据库
var usersDB = map[int]schema.UsersDBItem{}

// Get 获取用户信息
func (a *UsersService) Get(id int) (schema.UsersItem, error) {
	user, exists := usersDB[id]
	if !exists {
		return schema.UsersItem{}, errors.New("用户未找到")
	}

	// 类型转换
	return schema.UsersItem(user), nil
}

// Create 创建新用户
func (a *UsersService) Create(item *schema.UsersDBItem) error {
	// 模拟生成新用户ID
	newID := len(usersDB) + 1
	item.User_ID = newID
	usersDB[newID] = *item
	return nil
}

// Update 更新用户信息
func (a *UsersService) Update(id int, item *schema.UsersDBItem) error {
	_, exists := usersDB[id]
	if !exists {
		return errors.New("用户未找到")
	}

	// 更新用户信息
	item.User_ID = id
	usersDB[id] = *item
	return nil
}

// Delete 删除用户
func (a *UsersService) Delete(id int) error {
	_, exists := usersDB[id]
	if !exists {
		return errors.New("用户未找到")
	}

	// 删除用户
	delete(usersDB, id)
	return nil
}
