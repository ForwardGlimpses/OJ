package gormx

import (
	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/global"
)

// 通用分页查询函数
func GetPaginatedData(model interface{}, where interface{}, page, pageSize int, orderBy string) (interface{}, int64, error) {
	var total int64
	var result interface{}

	// 查询总记录数
	err := global.DB.Model(model).Where(where).Count(&total).Error
	if err != nil {
		return nil, 0, errors.InternalServer("Failed to count records: " + err.Error())
	}

	// 分页查询数据
	err = global.DB.Model(model).
		Where(where).
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Order(orderBy).
		Find(&result).Error
	if err != nil {
		return nil, 0, errors.InternalServer("Failed to retrieve data: " + err.Error())
	}

	return result, total, nil
}
