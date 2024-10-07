package service

import (
	"errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
)

type SolutionServiceInterface interface {
	Get(id int) (schema.SolutionItem, error)
	Create(item *schema.SolutionDBItem) error
	Update(id int, item *schema.SolutionDBItem) error
	Delete(id int) error
}

var SolutionServiceInstance SolutionServiceInterface = &SolutionService{}

type SolutionService struct{}

// 模拟数据库
var solutionDB = map[int]schema.SolutionDBItem{}

// Get 获取解决方案信息
func (a *SolutionService) Get(id int) (schema.SolutionItem, error) {
	item, exists := solutionDB[id]
	if !exists {
		return schema.SolutionItem{}, errors.New("解决方案未找到")
	}

	// 类型转换
	return schema.SolutionItem(item), nil
}

// Create 创建新的解决方案
func (a *SolutionService) Create(item *schema.SolutionDBItem) error {
	// 模拟生成新解决方案ID
	newID := len(solutionDB) + 1
	item.Solution_ID = newID
	solutionDB[newID] = *item
	return nil
}

// Update 更新解决方案信息
func (a *SolutionService) Update(id int, item *schema.SolutionDBItem) error {
	_, exists := solutionDB[id]
	if !exists {
		return errors.New("解决方案未找到")
	}

	// 更新解决方案信息
	item.Solution_ID = id
	solutionDB[id] = *item
	return nil
}

// Delete 删除解决方案
func (a *SolutionService) Delete(id int) error {
	_, exists := solutionDB[id]
	if !exists {
		return errors.New("解决方案未找到")
	}

	// 删除解决方案
	delete(solutionDB, id)
	return nil
}
