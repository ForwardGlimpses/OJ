package service

import (
	"context"

	"github.com/ForwardGlimpses/OJ/server/pkg/global"
	"github.com/ForwardGlimpses/OJ/server/pkg/gormx"
	"github.com/ForwardGlimpses/OJ/server/pkg/logs"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
)

type SourceCodeServiceInterface interface {
	Get(ctx context.Context, id int) (*schema.SourceCodeItem, error)
	Query(ctx context.Context, params schema.SourceCodeParams) (schema.SourceCodeItems, int64, error)
	Create(ctx context.Context, item *schema.SourceCodeItem) (int, error)
	Update(ctx context.Context, id int, item *schema.SourceCodeItem) error
	Delete(ctx context.Context, id int) error
}

var SourceCodeSvc SourceCodeServiceInterface = &SourceCodeService{}

type SourceCodeService struct{}

// Query根据条件和分页查询获取用户列表
func (a *SourceCodeService) Query(ctx context.Context, params schema.SourceCodeParams) (schema.SourceCodeItems, int64, error) {
	// 初始化查询
	query := global.DB.WithContext(ctx).Model(&schema.SourceCodeDBItem{})

	// 应用过滤条件
	if params.SolutionID != 0 {
		query.Where("solution_id = ?", params.SolutionID)
	}

	// 使用通用分页函数并指定返回类型
	sources, total, err := gormx.GetPaginatedData[schema.SourceCodeDBItem](query, params.P, "id ASC")
	if err != nil {
		logs.Error("Failed to query source codes:", err)
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
func (a *SourceCodeService) Get(ctx context.Context, id int) (*schema.SourceCodeItem, error) {
	db := global.DB.WithContext(ctx)
	item := &schema.SourceCodeDBItem{}
	err := db.Where("solution_id = ?", id).First(item).Error
	if err != nil {
		logs.Error("Failed to get source code with ID:", id, "Error:", err)
		return nil, err
	}
	return item.ToItem(), nil
}

// Create 创建源代码
func (a *SourceCodeService) Create(ctx context.Context, item *schema.SourceCodeItem) (int, error) {
	db := global.DB.WithContext(ctx)
	err := db.Create(item.ToDBItem()).Error
	if err != nil {
		logs.Error("Failed to create source code:", err)
		return 0, err
	}
	return item.ID, nil
}

// Update 更新源代码
func (a *SourceCodeService) Update(ctx context.Context, id int, item *schema.SourceCodeItem) error {
	db := global.DB.WithContext(ctx)
	dbItem := item.ToDBItem()
	err := db.Where("solution_id = ?", id).Updates(dbItem).Error
	if err != nil {
		logs.Error("Failed to update source code with ID:", id, "Error:", err)
		return err
	}
	return nil
}

// Delete 删除源代码
func (a *SourceCodeService) Delete(ctx context.Context, id int) error {
	db := global.DB.WithContext(ctx)
	err := db.Where("solution_id = ?", id).Delete(&schema.SourceCodeItem{}).Error
	if err != nil {
		logs.Error("Failed to delete source code with ID:", id, "Error:", err)
		return err
	}
	return nil
}
