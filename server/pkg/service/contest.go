package service

import (
	"context"
	"fmt"

	"github.com/ForwardGlimpses/OJ/server/pkg/global"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
)

type ContestServiceInterface interface {
	Query(params schema.ContestParams) (schema.ContestItems, error)
	Get(id int) (*schema.ContestItem, error)
	Create(item *schema.ContestItem) (int, error)
	Update(id int, item *schema.ContestItem) error
	Delete(id int) error
}

var ContestSvc ContestServiceInterface = &ContestService{}

type ContestService struct{}

// Query 获取比赛信息列表
func (a *ContestService) Query(params schema.ContestParams) (schema.ContestItems, error) {
	db := global.DB.WithContext(context.Background())
	if params.Title != "" {
		db.Where("title = ?", params.Title)
	}

	var items schema.ContestDBItems
	err := db.Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items.ToItems(), nil
}

// Get 获取比赛信息
func (a *ContestService) Get(id int) (*schema.ContestItem, error) {
	db := global.DB.WithContext(context.Background())
	//var item *schema.ContestDBItem
	item := &schema.ContestDBItem{}
	fmt.Println("idddd ", id)
	err := db.Where("id = ?", id).First(item).Error
	if err != nil {
		return nil, err
	}
	return item.ToItem(), nil
}

// Create 创建比赛
func (a *ContestService) Create(item *schema.ContestItem) (int, error) {
	db := global.DB.WithContext(context.Background())
	err := db.Create(item.ToDBItem()).Error
	if err != nil {
		return 0, err
	}
	return item.ID, nil
}

// Update 更新比赛
func (a *ContestService) Update(id int, item *schema.ContestItem) error {
	db := global.DB.WithContext(context.Background())
	err := db.Where("id = ?", id).Updates(item.ToDBItem()).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除比赛
func (a *ContestService) Delete(id int) error {
	db := global.DB.WithContext(context.Background())
	err := db.Where("id = ?", id).Delete(&schema.ContestDBItem{}).Error
	if err != nil {
		return err
	}
	return nil
}
