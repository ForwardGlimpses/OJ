package service

import (
	"context"

	"github.com/ForwardGlimpses/OJ/server/pkg/global"
	"github.com/ForwardGlimpses/OJ/server/pkg/gormx"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
)

type SolutionServiceInterface interface {
	Get(id int) (*schema.SolutionItem, error)
	Query(params schema.SolutionParams) (schema.SolutionItems, int64, error)
	Create(item *schema.SolutionItem) (int, error)
	Update(id int, item *schema.SolutionItem) error
	Delete(id int) error
}

var SolutionSvc SolutionServiceInterface = &SolutionService{}

type SolutionService struct{}

// Query根据条件和分页查询获取用户列表
func (a *SolutionService) Query(params schema.SolutionParams) (schema.SolutionItems, int64, error) {
	// 初始化查询
	query := global.DB.Model(&schema.SolutionDBItem{})

	// 应用过滤条件
	if params.UserID != "" {
		query = query.Where("user_id = ?", params.UserID)
	}

	// 使用通用分页函数并指定返回类型
	solutions, total, err := gormx.GetPaginatedData[schema.SolutionDBItem](query, params.P, "id ASC")
	if err != nil {
		return nil, 0, err
	}

	// 转换结果为返回的模型类型
	var items schema.SolutionItems
	for _, solution := range solutions {
		items = append(items, solution.ToItem())
	}

	return items, total, nil
}

// Get 获取解题方案
func (a *SolutionService) Get(id int) (*schema.SolutionItem, error) {
	db := global.DB.WithContext(context.Background())
	//var item *schema.SolutionDBItem
	item := &schema.SolutionDBItem{}
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
