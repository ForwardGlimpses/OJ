package service

import (
	"context"

	"github.com/ForwardGlimpses/OJ/server/pkg/global"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
)

type UsersServiceInterface interface {
	Get(id int) (*schema.UsersItem, error)
	Query(params schema.UsersParams) (schema.UsersItems, error)
	Create(item *schema.UsersItem) (int, error)
	Update(id int, item *schema.UsersItem) error
	Delete(id int) error
}

var UsersServiceInstance UsersServiceInterface = &UsersService{}

type UsersService struct{}

// Query 获取用户信息列表
func (a *UsersService) Query(params schema.UsersParams) (schema.UsersItems, error) {
	db := global.DB.WithContext(context.Background())
	if params.Email != "" {
		db.Where("email = ?", params.Email)
	}

	var items schema.UsersDBItems
	err := db.Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items.ToItems(), nil
}

// Get 获取用户信息
func (a *UsersService) Get(id int) (*schema.UsersItem, error) {
	db := global.DB.WithContext(context.Background())
	//var item *schema.UsersDBItem
	item := &schema.UsersDBItem{}
	err := db.Where("id = ?", id).First(item).Error
	if err != nil {
		return nil, err
	}
	return item.ToItem(), nil
}

// Create 创建用户
func (a *UsersService) Create(item *schema.UsersItem) (int, error) {
	db := global.DB.WithContext(context.Background())
	err := db.Create(item.ToDBItem()).Error
	if err != nil {
		return 0, err
	}
	return item.ID, nil
}

// Update 更新用户信息
func (a *UsersService) Update(id int, item *schema.UsersItem) error {
	db := global.DB.WithContext(context.Background())
	err := db.Where("user_id = ?", id).Updates(item.ToDBItem()).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除用户
func (a *UsersService) Delete(id int) error {
	db := global.DB.WithContext(context.Background())
	err := db.Where("user_id = ?", id).Delete(&schema.UsersDBItem{}).Error
	if err != nil {
		return err
	}
	return nil
}
