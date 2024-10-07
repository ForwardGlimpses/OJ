package service

import (
	"errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
)

type SourceCodeServiceInterface interface {
	Get(id int) (schema.Source_codeItem, error)
	Create(item *schema.Source_codeDBItem) error
	Update(id int, item *schema.Source_codeDBItem) error
	Delete(id int) error
}

var SourceCodeServiceInstance SourceCodeServiceInterface = &SourceCodeService{}

type SourceCodeService struct{}

// 模拟数据库
var sourceCodeDB = map[int]schema.Source_codeDBItem{}

// Get 获取源代码信息
func (a *SourceCodeService) Get(id int) (schema.Source_codeItem, error) {
	item, exists := sourceCodeDB[id]
	if !exists {
		return schema.Source_codeItem{}, errors.New("源代码未找到")
	}

	// 类型转换
	return schema.Source_codeItem(item), nil
}

// Create 创建新的源代码
func (a *SourceCodeService) Create(item *schema.Source_codeDBItem) error {
	// 模拟生成新源代码ID
	newID := len(sourceCodeDB) + 1
	item.Solution_ID = newID
	sourceCodeDB[newID] = *item
	return nil
}

// Update 更新源代码信息
func (a *SourceCodeService) Update(id int, item *schema.Source_codeDBItem) error {
	_, exists := sourceCodeDB[id]
	if !exists {
		return errors.New("源代码未找到")
	}

	// 更新源代码信息
	item.Solution_ID = id
	sourceCodeDB[id] = *item
	return nil
}

// Delete 删除源代码
func (a *SourceCodeService) Delete(id int) error {
	_, exists := sourceCodeDB[id]
	if !exists {
		return errors.New("源代码未找到")
	}

	// 删除源代码
	delete(sourceCodeDB, id)
	return nil
}
