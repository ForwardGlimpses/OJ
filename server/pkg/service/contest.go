package service

import (
	"context"
	"fmt"

	"github.com/ForwardGlimpses/OJ/server/pkg/global"
	"github.com/ForwardGlimpses/OJ/server/pkg/gormx"
	"github.com/ForwardGlimpses/OJ/server/pkg/logs"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
)

type ContestServiceInterface interface {
	Query(params schema.ContestParams) (schema.ContestItems, int64, error)
	Get(id int) (*schema.ContestItem, error)
	Create(item *schema.ContestItem) (int, error)
	Update(id int, item *schema.ContestItem) error
	Delete(id int) error
}

var ContestSvc ContestServiceInterface = &ContestService{}

type ContestService struct{}

// Query根据条件和分页查询获取用户列表
func (a *ContestService) Query(params schema.ContestParams) (schema.ContestItems, int64, error) {
	// 初始化查询
	query := global.DB.Model(&schema.ContestDBItem{})

	// 应用过滤条件
	if params.Title != "" {
		query = query.Where("title = ?", params.Title)
	}

	// 使用通用分页函数并指定返回类型
	contests, total, err := gormx.GetPaginatedData[schema.ContestDBItem](query, params.P, "id ASC")
	if err != nil {
		logs.Error("Failed to query contests:", err)
		return nil, 0, err
	}

	// 转换结果为返回的模型类型
	var items schema.ContestItems
	for _, contest := range contests {
		items = append(items, contest.ToItem())
	}

	return items, total, nil
}

// Get 获取比赛信息
func (a *ContestService) Get(id int) (*schema.ContestItem, error) {
	db := global.DB.WithContext(context.Background())
	//var item *schema.ContestDBItem
	item := &schema.ContestDBItem{}
	fmt.Println("idddd ", id)
	err := db.Where("id = ?", id).First(item).Error
	if err != nil {
		logs.Error("Failed to get contest with ID:", id, "Error:", err)
		return nil, err
	}
	return item.ToItem(), nil
}

// Create 创建比赛
func (a *ContestService) Create(item *schema.ContestItem) (int, error) {
	db := global.DB.WithContext(context.Background())
	dbItem := item.ToDBItem()
	err := db.Create(dbItem).Error
	if err != nil {
		logs.Error("Failed to create contest:", err)
		return 0, err
	}
	return dbItem.ID, nil
}

// Update 更新比赛
func (a *ContestService) Update(id int, item *schema.ContestItem) error {
	db := global.DB.WithContext(context.Background())
	dbItem := item.ToDBItem()
	err := db.Where("id = ?", id).Updates(dbItem).Error
	if err != nil {
		logs.Error("Failed to update contest with ID:", id, "Error:", err)
		return err
	}
	return nil
}

// Delete 删除比赛
func (a *ContestService) Delete(id int) error {
	db := global.DB.WithContext(context.Background())
	err := db.Where("id = ?", id).Delete(&schema.ContestDBItem{}).Error
	if err != nil {
		logs.Error("Failed to delete contest with ID:", id, "Error:", err)
		return err
	}
	return nil
}
