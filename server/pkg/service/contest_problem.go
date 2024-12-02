package service

import (
	"context"

	"github.com/ForwardGlimpses/OJ/server/pkg/global"
	"github.com/ForwardGlimpses/OJ/server/pkg/gormx"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
)

type ContestProblemServiceInterface interface {
	Query(params schema.ContestProblemParams) (schema.ContestProblemItems, int64, error)
	Get(id int) (*schema.ContestProblemItem, error)
	Create(item *schema.ContestProblemItem) (int, error)
	Update(id int, item *schema.ContestProblemItem) error
	Delete(id int) error
}

var ContestProblemSvc ContestProblemServiceInterface = &ContestProblemService{}

type ContestProblemService struct{}

// Query根据条件和分页查询获取用户列表
func (a *ContestProblemService) Query(params schema.ContestProblemParams) (schema.ContestProblemItems, int64, error) {
	// 初始化查询
	query := global.DB.Model(&schema.ContestProblemDBItem{})

	// 应用过滤条件
	if params.Title != "" {
		query = query.Where("title = ?", params.Title)
	}

	// 使用通用分页函数并指定返回类型
	contestproblems, total, err := gormx.GetPaginatedData[schema.ContestProblemDBItem](query, params.P, "id ASC")
	if err != nil {
		return nil, 0, err
	}

	// 转换结果为返回的模型类型
	var items schema.ContestProblemItems
	for _, contestproblem := range contestproblems {
		items = append(items, contestproblem.ToItem())
	}

	return items, total, nil
}

// Get 获取比赛问题信息
func (a *ContestProblemService) Get(id int) (*schema.ContestProblemItem, error) {
	db := global.DB.WithContext(context.Background())
	//var item *schema.ContestProblemDBItem
	item := &schema.ContestProblemDBItem{}
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
