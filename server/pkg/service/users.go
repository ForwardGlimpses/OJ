package service

import (
	"context"
	"github.com/ForwardGlimpses/OJ/server/pkg/errors"
	"github.com/ForwardGlimpses/OJ/server/pkg/global"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
	"github.com/ForwardGlimpses/OJ/server/pkg/gormx"
	"golang.org/x/crypto/bcrypt"
)

type UsersServiceInterface interface {
	Get(id int) (*schema.UsersItem, error)
	GetWithEmail(email string) (*schema.UsersItem, error)
	Query(params schema.UsersParams) (schema.UsersItems, int64, error)
	Create(item *schema.UsersItem) (int, error)
	Update(id int, item *schema.UsersItem) error
	Delete(id int) error
	Register(name, email, password, school string) (*schema.UsersItem, error) // Register方法只需验证、处理和调用Create
}

var UserSvc UsersServiceInterface = &UsersService{}

type UsersService struct{}

// Query根据条件和分页查询获取用户列表
func (a *UsersService) Query(params schema.UsersParams) (schema.UsersItems, int64, error) {

	// 构造查询条件
	where := map[string]interface{}{}
	if params.Email != "" {
		where["email"] = params.Email
	}
	if params.Name != "" {
		where["name"] = params.Name
	}
	if params.School != "" {
		where["school"] = params.School
	}

	// 获取符合条件的分页数据和总数
	users, total, err := gormx.GetPaginatedData(&schema.UsersDBItem{}, where, params.Page, params.PageSize, "id ASC")
	if err != nil {
		return nil, 0, err
	}

	// 转换为返回的模型
	var items schema.UsersItems
	for _, user := range users.([]*schema.UsersDBItem) {
		items = append(items, user.ToItem())
	}

	return items, total, nil
}

// Get 获取用户信息
func (a *UsersService) Get(id int) (*schema.UsersItem, error) {
	db := global.DB.WithContext(context.Background())
	item := &schema.UsersDBItem{}
	err := db.Where("id = ?", id).First(item).Error
	if err != nil {
		return nil, err
	}
	return item.ToItem(), nil
}

// GetWithEmail 根据邮箱获取用户信息
func (a *UsersService) GetWithEmail(email string) (*schema.UsersItem, error) {
	db := global.DB.WithContext(context.Background())
	item := &schema.UsersDBItem{}
	err := db.Where("email = ?", email).First(item).Error
	if err != nil {
		return nil, err
	}
	return item.ToItem(), nil
}

// Create 创建用户
func (a *UsersService) Create(item *schema.UsersItem) (int, error) {
	db := global.DB.WithContext(context.Background())
	err := db.Create(item.ToDBItem()).Error
	if err != nil {
		return 0, err
	}
	return item.ID, nil
}

// Update 更新用户信息
func (a *UsersService) Update(id int, item *schema.UsersItem) error {
	db := global.DB.WithContext(context.Background())
	err := db.Where("id = ?", id).Updates(item.ToDBItem()).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除用户
func (a *UsersService) Delete(id int) error {
	db := global.DB.WithContext(context.Background())
	err := db.Where("id = ?", id).Delete(&schema.UsersDBItem{}).Error
	if err != nil {
		return err
	}
	return nil
}

// Register 用户注册方法
func (a *UsersService) Register(name, email, password, school string) (*schema.UsersItem, error) {
	// 检查邮箱是否已经注册
	_, err := a.GetWithEmail(email)
	if err == nil {
		// 如果用户已存在，返回错误
		return nil, errors.InvalidInput("email already registered")
	}

	// 哈希密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.InternalServer("failed to hash password")
	}

	// 创建新的用户实例
	newUser := &schema.UsersItem{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Submit:   0,
		Solved:   0,
		School:   school,
	}
	// 调用 Create 方法将用户保存到数据库
	_, err = a.Create(newUser)
	if err != nil {
		return nil, errors.InternalServer("failed to create user")
	}
	return newUser, nil
}
