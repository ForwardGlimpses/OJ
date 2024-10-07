package service

import (
	"errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
)

type ProblemServiceInterface interface {
	Get(id int) (schema.ProblemItem, error)
	Create(item *schema.ProblemDBItem) error
	Update(id int, item *schema.ProblemDBItem) error
	Delete(id int) error
}

var ProblemServiceInstance ProblemServiceInterface = &ProblemService{}

type ProblemService struct{}

// 模拟数据库
var problemDB = map[int]schema.ProblemDBItem{}

// Get 获取问题
func (a *ProblemService) Get(id int) (schema.ProblemItem, error) {
	problem, exists := problemDB[id]
	if !exists {
		return schema.ProblemItem{}, errors.New("问题未找到")
	}
	return schema.ProblemItem(problem), nil
}

// Create 创建新问题
func (a *ProblemService) Create(item *schema.ProblemDBItem) error {
	// 假设数据库ID是递增的，这里简单模拟一下
	newID := len(problemDB) + 1
	item.ID = newID
	problemDB[newID] = *item
	return nil
}

// Update 更新指定ID的问题
func (a *ProblemService) Update(id int, item *schema.ProblemDBItem) error {
	_, exists := problemDB[id]
	if !exists {
		return errors.New("问题未找到")
	}
	// 更新问题
	item.ID = id
	problemDB[id] = *item
	return nil
}

// Delete 删除指定ID的问题
func (a *ProblemService) Delete(id int) error {
	_, exists := problemDB[id]
	if !exists {
		return errors.New("问题未找到")
	}
	// 从数据库删除问题
	delete(problemDB, id)
	return nil
}
