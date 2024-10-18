package service

import (
	"context"

	"github.com/ForwardGlimpses/OJ/server/pkg/global"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
)

type ProblemServiceInterface interface {
	Query(params schema.ProblemParams) (schema.ProblemItems, error)
	Get(id int) (*schema.ProblemItem, error)
	Create(item *schema.ProblemItem) (int, error)
	Update(id int, item *schema.ProblemItem) error
	Delete(id int) error
}

var ProblemServiceInstance ProblemServiceInterface = &ProblemService{}

type ProblemService struct{}

// Query 获取比赛信息列表
func (a *ProblemService) Query(params schema.ProblemParams) (schema.ProblemItems, error) {
	db := global.DB.WithContext(context.Background())
	if params.Title != "" {
		db.Where("title = ?", params.Title)
	}

	var items schema.ProblemDBItems
	err := db.Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items.ToItems(), nil
}

// Get 通过ID从数据库获取题目
func (a *ProblemService) Get(id int) (*schema.ProblemItem, error) {
	db := global.DB.WithContext(context.Background())
	var item *schema.ProblemDBItem
	err := db.Where("id = ?", id).First(item).Error
	if err != nil {
		return nil, err
	}
	return item.ToItem(), nil
}

// Create 将 ProblemItem 转换为 ProblemDBItem 并存入数据库
func (a *ProblemService) Create(item *schema.ProblemItem) (int, error) {
	db := global.DB.WithContext(context.Background())
	err := db.Create(item.ToDBItem()).Error
	if err != nil {
		return 0, err
	}
	return item.ID, nil
}

// Update 更新题目信息
func (a *ProblemService) Update(id int, item *schema.ProblemItem) error {
	db := global.DB.WithContext(context.Background())
	err := db.Where("id = ?", id).Updates(item.ToDBItem()).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete 根据ID删除题目
func (a *ProblemService) Delete(id int) error {
	db := global.DB.WithContext(context.Background())
	err := db.Where("id = ?", id).Delete(&schema.ProblemDBItem{}).Error
	if err != nil {
		return err
	}
	return nil
}
