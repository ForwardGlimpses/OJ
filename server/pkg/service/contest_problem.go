package service

import (
	"errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
)

type ContestProblemServiceInterface interface {
	Get(id int) (schema.Contest_ProblemItem, error)
	Create(item *schema.Contest_ProblemDBItem) error
	Update(id int, item *schema.Contest_ProblemDBItem) error
	Delete(id int) error
}

var ContestProblemServiceInstance ContestProblemServiceInterface = &ContestProblemService{}

type ContestProblemService struct{}

// 模拟数据库
var contestProblemDB = map[int]schema.Contest_ProblemDBItem{}

// Get 获取比赛问题信息
func (a *ContestProblemService) Get(id int) (schema.Contest_ProblemItem, error) {
	item, exists := contestProblemDB[id]
	if !exists {
		return schema.Contest_ProblemItem{}, errors.New("比赛问题未找到")
	}

	// 类型转换
	return schema.Contest_ProblemItem(item), nil
}

// Create 创建新的比赛问题
func (a *ContestProblemService) Create(item *schema.Contest_ProblemDBItem) error {
	// 模拟生成新比赛问题ID
	newID := len(contestProblemDB) + 1
	item.Problem_ID = newID
	contestProblemDB[newID] = *item
	return nil
}

// Update 更新比赛问题信息
func (a *ContestProblemService) Update(id int, item *schema.Contest_ProblemDBItem) error {
	_, exists := contestProblemDB[id]
	if !exists {
		return errors.New("比赛问题未找到")
	}

	// 更新比赛问题信息
	item.Problem_ID = id
	contestProblemDB[id] = *item
	return nil
}

// Delete 删除比赛问题
func (a *ContestProblemService) Delete(id int) error {
	_, exists := contestProblemDB[id]
	if !exists {
		return errors.New("比赛问题未找到")
	}

	// 删除比赛问题
	delete(contestProblemDB, id)
	return nil
}
