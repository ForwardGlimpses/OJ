package service

import (
	"context"

	"github.com/ForwardGlimpses/OJ/server/pkg/global"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
)

type SourceCodeServiceInterface interface {
	Get(id int) (*schema.Source_codeItem, error)
	Query(params schema.SourceCodeParams) (schema.SourceCodeItems, error)
	Create(item *schema.Source_codeItem) (int, error)
	Update(id int, item *schema.Source_codeItem) error
	Delete(id int) error
}

var SourceCodeServiceInstance SourceCodeServiceInterface = &SourceCodeService{}

type SourceCodeService struct{}

// Query 获取源代码列表
func (a *SourceCodeService) Query(params schema.SourceCodeParams) (schema.SourceCodeItems, error) {
	db := global.DB.WithContext(context.Background())
	if params.SolutionID != 0 {
		db.Where("solution_id = ?", params.SolutionID)
	}

	var items schema.Source_codeItems
	err := db.Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items.ToItems(), nil
}

// Get 获取源代码
func (a *SourceCodeService) Get(id int) (*schema.Source_codeItem, error) {
	db := global.DB.WithContext(context.Background())
	var item *schema.Source_codeItem
	err := db.Where("solution_id = ?", id).First(item).Error
	if err != nil {
		return nil, err
	}
	return item.ToItem(), nil
}

// Create 创建源代码
func (a *SourceCodeService) Create(item *schema.Source_codeItem) (int, error) {
	db := global.DB.WithContext(context.Background())
	err := db.Create(item.ToDBItem()).Error
	if err != nil {
		return 0, err
	}
	return item.ID, nil
}

// Update 更新源代码
func (a *SourceCodeService) Update(id int, item *schema.Source_codeItem) error {
	db := global.DB.WithContext(context.Background())
	err := db.Where("solution_id = ?", id).Updates(item.ToDBItem()).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除源代码
func (a *SourceCodeService) Delete(id int) error {
	db := global.DB.WithContext(context.Background())
	err := db.Where("solution_id = ?", id).Delete(&schema.Source_codeItem{}).Error
	if err != nil {
		return err
	}
	return nil
}
