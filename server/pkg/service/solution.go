package service

import (
	"context"

	"github.com/ForwardGlimpses/OJ/server/pkg/global"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
)

type SolutionServiceInterface interface {
	Get(id int) (*schema.SolutionItem, error)
	Query(params schema.SolutionParams) (schema.SolutionItems, error)
	Create(item *schema.SolutionItem) (int, error)
	Update(id int, item *schema.SolutionItem) error
	Delete(id int) error
}

var SolutionServiceInstance SolutionServiceInterface = &SolutionService{}

type SolutionService struct{}

// Query 获取解题方案列表
func (a *SolutionService) Query(params schema.SolutionParams) (schema.SolutionItems, error) {
	db := global.DB.WithContext(context.Background())
	if params.UserID != "" {
		db.Where("user_id = ?", params.UserID)
	}

	var items schema.SolutionDBItems
	err := db.Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items.ToItems(), nil
}

// Get 获取解题方案
func (a *SolutionService) Get(id int) (*schema.SolutionItem, error) {
	db := global.DB.WithContext(context.Background())
	var item *schema.SolutionDBItem
	err := db.Where("id = ?", id).First(item).Error
	if err != nil {
		return nil, err
	}
	return item.ToItem(), nil
}

// Create 创建解题方案
func (a *SolutionService) Create(item *schema.SolutionItem) (int, error) {
	db := global.DB.WithContext(context.Background())
	err := db.Create(item.ToDBItem()).Error
	if err != nil {
		return 0, err
	}
	return item.ID, nil
}

// Update 更新解题方案
func (a *SolutionService) Update(id int, item *schema.SolutionItem) error {
	db := global.DB.WithContext(context.Background())
	err := db.Where("id = ?", id).Updates(item.ToDBItem()).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除解题方案
func (a *SolutionService) Delete(id int) error {
	db := global.DB.WithContext(context.Background())
	err := db.Where("id = ?", id).Delete(&schema.SolutionDBItem{}).Error
	if err != nil {
		return err
	}
	return nil
}
