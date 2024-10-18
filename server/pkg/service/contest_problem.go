package service

import (
	"context"

	"github.com/ForwardGlimpses/OJ/server/pkg/global"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
)

type ContestProblemServiceInterface interface {
	Get(id int) (*schema.ContestProblemItem, error)
	Create(item *schema.ContestProblemItem) (int, error)
	Update(id int, item *schema.ContestProblemItem) error
	Delete(id int) error
}

var ContestProblemServiceInstance ContestProblemServiceInterface = &ContestProblemService{}

type ContestProblemService struct{}

// Query 获取比赛问题信息列表
func (a *ContestProblemService) Query(params schema.ContestProblemParams) (schema.ContestProblemItems, error) {
	db := global.DB.WithContext(context.Background())
	if params.Title != "" {
		db.Where("title = ?", params.Title)
	}

	var items schema.ContestProblemDBItems
	err := db.Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items.ToItems(), nil
}

// Get 获取比赛问题信息
func (a *ContestProblemService) Get(id int) (*schema.ContestProblemItem, error) {
	db := global.DB.WithContext(context.Background())
	var item *schema.ContestProblemDBItem
	err := db.Where("id = ?", id).First(item).Error
	if err != nil {
		return nil, err
	}

	return item.ToItem(), nil
}

// Create 创建新的比赛问题
func (a *ContestProblemService) Create(item *schema.ContestProblemItem) (int, error) {
	db := global.DB.WithContext(context.Background())
	err := db.Create(item.ToDBItem()).Error
	if err != nil {
		return 0, err
	}
	return item.ID, nil
}

// Update 更新比赛问题信息
func (a *ContestProblemService) Update(id int, item *schema.ContestProblemItem) error {
	db := global.DB.WithContext(context.Background())
	err := db.Where("id = ?", id).Updates(item.ToDBItem()).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除比赛问题
func (a *ContestProblemService) Delete(id int) error {
	db := global.DB.WithContext(context.Background())
	err := db.Where("id = ?", id).Delete(&schema.ContestProblemDBItem{}).Error
	if err != nil {
		return err
	}
	return nil
}