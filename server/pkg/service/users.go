package service

import (
	"context"
	"time"

	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/global"
	"github.com/ForwardGlimpses/OJ/server/pkg/gormx"
	"github.com/ForwardGlimpses/OJ/server/pkg/logs"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
	"golang.org/x/crypto/bcrypt"
)

type UsersServiceInterface interface {
	Get(ctx context.Context, id int) (*schema.UsersItem, error)
	GetWithEmail(ctx context.Context, email string) (*schema.UsersItem, error)
	Query(ctx context.Context, params schema.UsersParams) (schema.UsersItems, int64, error)
	Update(ctx context.Context, id int, item *schema.UsersItem) error
	Delete(ctx context.Context, id int) error
	Create(ctx context.Context, item schema.UsersItem) (int, error) // Register方法只需验证、处理和调用Create
}

var UserSvc UsersServiceInterface = &UsersService{}

type UsersService struct{}

// Query根据条件和分页查询获取用户列表
func (a *UsersService) Query(ctx context.Context, params schema.UsersParams) (schema.UsersItems, int64, error) {
	// 初始化查询
	query := global.DB.WithContext(ctx).Model(&schema.UsersDBItem{})

	// 应用过滤条件
	if params.Email != "" {
		query = query.Where("email = ?", params.Email)
	}
	if params.Name != "" {
		query = query.Where("name = ?", params.Name)
	}
	if params.School != "" {
		query = query.Where("school = ?", params.School)
	}

	// 使用通用分页函数并指定返回类型
	users, total, err := gormx.GetPaginatedData[schema.UsersDBItem](query, params.P, "id ASC")
	if err != nil {
		logs.Error("Failed to query users:", err)
		return nil, 0, err
	}

	// 转换结果为返回的模型类型
	var items schema.UsersItems
	for _, user := range users {
		items = append(items, user.ToItem())
	}

	return items, total, nil
}

// Get 获取用户信息
func (a *UsersService) Get(ctx context.Context, id int) (*schema.UsersItem, error) {
	db := global.DB.WithContext(ctx)
	item := &schema.UsersDBItem{}
	err := db.Where("id = ?", id).First(item).Error
	if err != nil {
		logs.Error("Failed to get user with ID:", id, "Error:", err)
		return nil, err
	}
	return item.ToItem(), nil
}

// GetWithEmail 根据邮箱获取用户信息
func (a *UsersService) GetWithEmail(ctx context.Context, email string) (*schema.UsersItem, error) {
	db := global.DB.WithContext(ctx)
	item := &schema.UsersDBItem{}
	err := db.Where("email = ?", email).First(item).Error
	if err != nil {
		logs.Error("Failed to get user with email:", email, "Error:", err)
		return nil, err
	}
	return item.ToItem(), nil
}

// 用户注册方法
func (a *UsersService) Create(ctx context.Context, item schema.UsersItem) (int, error) {
	// 哈希密码

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(item.Password), bcrypt.DefaultCost)
	if err != nil {
		logs.Error("Failed to hash password:", err)
		return 0, errors.InternalServer("failed to hash password")
	}

	item.Password = string(hashedPassword)
	item.Accesstime = time.Now()

	dbItem := item.ToDBItem()
	// 将用户数据保存到数据库
	err = global.DB.WithContext(ctx).Create(dbItem).Error
	if err != nil {
		logs.Error("Failed to create user:", err)
		return 0, errors.InternalServer("failed to create user")
	}

	return dbItem.ID, nil
}

// Update 更新用户信息
func (a *UsersService) Update(ctx context.Context, id int, item *schema.UsersItem) error {
	db := global.DB.WithContext(ctx)
	dbItem := item.ToDBItem()
	err := db.Where("id = ?", id).Updates(dbItem).Error
	if err != nil {
		logs.Error("Failed to update user with ID:", id, "Error:", err)
		return err
	}
	return nil
}

// Delete 删除用户
func (a *UsersService) Delete(ctx context.Context, id int) error {
	db := global.DB.WithContext(ctx)
	err := db.Where("id = ?", id).Delete(&schema.UsersDBItem{}).Error
	if err != nil {
		logs.Error("Failed to delete user with ID:", id, "Error:", err)
		return err
	}
	return nil
}
