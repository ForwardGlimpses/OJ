package service

import (
	"context"
	"fmt"

	"github.com/ForwardGlimpses/OJ/server/pkg/global"
	"github.com/ForwardGlimpses/OJ/server/pkg/gormx"
	"github.com/ForwardGlimpses/OJ/server/pkg/logs"
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
		logs.Error("Failed to query solutions:", err)
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
		logs.Error("Failed to get solution with ID:", id, "Error:", err)
		return nil, err
	}
	return item.ToItem(), nil
}

// Create 创建解题方案
func (a *SolutionService) Create(item *schema.SolutionItem) (int, error) {
	db := global.DB.WithContext(context.Background())
	err := db.Create(item.ToDBItem()).Error
	if err != nil {
		logs.Error("Failed to create solution:", err)
		return 0, err
	}
	// 手动查询填充 ID
	var createdItem schema.SolutionDBItem
	err = db.Where("problem_id = ? AND user_id = ? AND status = ?", item.ProblemID, item.UserID, "Pending").First(&createdItem).Error
	if err != nil {
		logs.Error("Failed to get created solution:", err)
		return 0, err
	}

	// 更新 SolutionItem 的 ID
	item.ID = createdItem.ID
	fmt.Println("Solution created with ID:", item.ID)
	return item.ID, nil
}

// Update 更新解题方案
func (a *SolutionService) Update(id int, item *schema.SolutionItem) error {
	db := global.DB.WithContext(context.Background())
	err := db.Where("id = ?", id).Updates(item.ToDBItem()).Error
	if err != nil {
		logs.Error("Failed to update solution with ID:", id, "Error:", err)
		return err
	}
	return nil
}

// Delete 删除解题方案
func (a *SolutionService) Delete(id int) error {
	db := global.DB.WithContext(context.Background())
	err := db.Where("id = ?", id).Delete(&schema.SolutionDBItem{}).Error
	if err != nil {
		logs.Error("Failed to delete solution with ID:", id, "Error:", err)
		return err
	}
	return nil
}
