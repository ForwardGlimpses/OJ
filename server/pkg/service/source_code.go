package service

import (
	"context"

	"github.com/ForwardGlimpses/OJ/server/pkg/global"
	"github.com/ForwardGlimpses/OJ/server/pkg/gormx"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
)

type SourceCodeServiceInterface interface {
	Get(id int) (*schema.SourceCodeItem, error)
	Query(params schema.SourceCodeParams) (schema.SourceCodeItems, int64, error)
	Create(item *schema.SourceCodeItem) (int, error)
	Update(id int, item *schema.SourceCodeItem) error
	Delete(id int) error
}

var SourceCodeSvc SourceCodeServiceInterface = &SourceCodeService{}

type SourceCodeService struct{}

// Query根据条件和分页查询获取用户列表
func (a *SourceCodeService) Query(params schema.SourceCodeParams) (schema.SourceCodeItems, int64, error) {
	// 初始化查询
	query := global.DB.Model(&schema.SourceCodeDBItem{})

	// 应用过滤条件
	if params.SolutionID != 0 {
		query.Where("solution_id = ?", params.SolutionID)
	}

	// 使用通用分页函数并指定返回类型
	sources, total, err := gormx.GetPaginatedData[schema.SourceCodeDBItem](query, params.P, "id ASC")
	if err != nil {
		return nil, 0, err
	}

	// 转换结果为返回的模型类型
	var items schema.SourceCodeItems
	for _, source := range sources {
		items = append(items, source.ToItem())
	}

	return items, total, nil
}

// Get 获取源代码
func (a *SourceCodeService) Get(id int) (*schema.SourceCodeItem, error) {
	db := global.DB.WithContext(context.Background())
	item := &schema.SourceCodeDBItem{}
	err := db.Where("solution_id = ?", id).First(item).Error
	if err != nil {
		return nil, err
	}
	return item.ToItem(), nil
}

// Create 创建源代码
func (a *SourceCodeService) Create(item *schema.SourceCodeItem) (int, error) {
	db := global.DB.WithContext(context.Background())
	err := db.Create(item.ToDBItem()).Error
	if err != nil {
		return 0, err
	}
	return item.ID, nil
}

// Update 更新源代码
func (a *SourceCodeService) Update(id int, item *schema.SourceCodeItem) error {
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
	err := db.Where("solution_id = ?", id).Delete(&schema.SourceCodeItem{}).Error
	if err != nil {
		return err
	}
	return nil
}
