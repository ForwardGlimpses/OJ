package gormx

import (
	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
	"gorm.io/gorm"
)

// 通用分页查询函数，支持返回指定类型的数据
func GetPaginatedData[T any](db *gorm.DB, p schema.P, orderBy string) ([]T, int64, error) {
	var total int64
	var result []T
	var t T

	// 查询总记录数
	err := db.Model(&t).Count(&total).Error
	if err != nil {
		return nil, 0, errors.InternalServer("查询数据失败: %v", err)
	}

	// 分页查询数据
	err = db.
		Limit(p.PageSize).
		Offset((p.Page - 1) * p.PageSize).
		Order(orderBy).
		Find(&result).Error
	if err != nil {
		return nil, 0, errors.InternalServer("分页查询数据失败: %v", err)
	}

	return result, total, nil
}
