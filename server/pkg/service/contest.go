package service

import (
	"errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
)

type ContestServiceInterface interface {
	Get(id int) (schema.ContestItem, error)
	Create(item *schema.ContestDBItem) error
	Update(id int, item *schema.ContestDBItem) error
	Delete(id int) error
}

var ContestServiceInstance ContestServiceInterface = &ContestService{}

type ContestService struct{}

// 模拟数据库
var contestDB = map[int]schema.ContestDBItem{}

// Get 获取比赛信息
func (a *ContestService) Get(id int) (schema.ContestItem, error) {
	item, exists := contestDB[id]
	if !exists {
		return schema.ContestItem{}, errors.New("比赛未找到")
	}

	// 类型转换
	return schema.ContestItem(item), nil
}

// Create 创建新的比赛
func (a *ContestService) Create(item *schema.ContestDBItem) error {
	// 模拟生成新比赛ID
	newID := len(contestDB) + 1
	item.Contest_ID = newID
	contestDB[newID] = *item
	return nil
}

// Update 更新比赛信息
func (a *ContestService) Update(id int, item *schema.ContestDBItem) error {
	_, exists := contestDB[id]
	if !exists {
		return errors.New("比赛未找到")
	}

	// 更新比赛信息
	item.Contest_ID = id
	contestDB[id] = *item
	return nil
}

// Delete 删除比赛
func (a *ContestService) Delete(id int) error {
	_, exists := contestDB[id]
	if !exists {
		return errors.New("比赛未找到")
	}

	// 删除比赛
	delete(contestDB, id)
	return nil
}
