package service

import (
	"context" // 导入 json 包
	"fmt"
	"time"

	"github.com/ForwardGlimpses/OJ/server/pkg/global"
	"github.com/ForwardGlimpses/OJ/server/pkg/gormx"
	"github.com/ForwardGlimpses/OJ/server/pkg/judge"
	"github.com/ForwardGlimpses/OJ/server/pkg/logs"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
	//"github.com/google/uuid"
)

type ProblemServiceInterface interface {
	Query(params schema.ProblemParams) (schema.ProblemItems, int64, error)
	Get(id int) (*schema.ProblemItem, error)
	Create(item *schema.ProblemItem) (int, error)
	Update(id int, item *schema.ProblemItem) error
	Delete(id int) error
	//TODO 提交 判断 存储提交结果 返回提交结果id给前端 前端拿提交的id来查询status
	Submit(id int, userId int, inputCode string) (int, error)
}

var ProblemSvc ProblemServiceInterface = &ProblemService{}

type ProblemService struct{}

// Query根据条件和分页查询获取用户列表
func (a *ProblemService) Query(params schema.ProblemParams) (schema.ProblemItems, int64, error) {
	// 初始化查询
	query := global.DB.Model(&schema.ProblemDBItem{})

	// 应用过滤条件
	if params.ProblemID != 0 {
		query = query.Where("Problem_id = ?", params.ProblemID)
	}

	// 使用通用分页函数并指定返回类型
	problems, total, err := gormx.GetPaginatedData[schema.ProblemDBItem](query, params.P, "id ASC")
	if err != nil {
		logs.Error("Failed to query problems:", err)
		return nil, 0, err
	}

	// 转换结果为返回的模型类型
	var items schema.ProblemItems
	for _, problem := range problems {
		items = append(items, problem.ToItem())
	}

	return items, total, nil
}

// Get 通过ID从数据库获取题目
func (a *ProblemService) Get(id int) (*schema.ProblemItem, error) {
	db := global.DB.WithContext(context.Background())
	item := &schema.ProblemDBItem{}
	err := db.Where("id = ?", id).First(item).Error
	if err != nil {
		logs.Error("Failed to get problem with ID:", id, "Error:", err)
		return nil, err
	}
	logs.Info("Successfully retrieved problem with ID:", id)
	return item.ToItem(), nil
}

// Create 将 ProblemItem 转换为 ProblemDBItem 并存入数据库
func (a *ProblemService) Create(item *schema.ProblemItem) (int, error) {
	db := global.DB.WithContext(context.Background())
	dbItem := item.ToDBItem()
	err := db.Create(dbItem).Error
	if err != nil {
		logs.Error("Failed to create problem:", err)
		return 0, err
	}
	return dbItem.ID, nil
}

// Update 更新题目信息
func (a *ProblemService) Update(id int, item *schema.ProblemItem) error {
	db := global.DB.WithContext(context.Background())
	err := db.Where("id = ?", id).Updates(item.ToDBItem()).Error
	if err != nil {
		logs.Error("Failed to update problem with ID:", id, "Error:", err)
		return err
	}
	return nil
}

// Delete 根据ID删除题目
func (a *ProblemService) Delete(id int) error {
	db := global.DB.WithContext(context.Background())
	err := db.Where("id = ?", id).Delete(&schema.ProblemDBItem{})
	if err.Error != nil {
		logs.Error("Failed to delete problem with ID:", id, "Error:", err.Error)
		return err.Error
	}
	if err.RowsAffected == 0 {
		return fmt.Errorf("no record found with id %d", id)
	}
	return nil
}

func (a *ProblemService) Submit(id int, userId int, code string) (int, error) {

	//db := global.DB.WithContext(context.Background())

	// Step 1: 获取题目详细信息
	problem, err := a.Get(id)
	if err != nil {
		logs.Error("Failed to get problem with ID:", id, "Error:", err)
		return 0, err
	}

	submission := &schema.SolutionItem{
		ProblemID: id,
		UserID:    userId,
		Status:    "Pending",
	}

	submissionId, err := SolutionSvc.Create(submission)
	if err != nil {
		logs.Error("Failed to create submission:", err)
		return 0, err
	}

	resp, err := judge.Submit(judge.Request{
		ID:     submissionId,
		Code:   code,
		Input:  problem.Input,
		Output: problem.Output,
	})

	if err != nil {
		logs.Error("Failed to judge:", err)
		return 0, err
	}

	// Step 6: 将评测结果存储到数据库
	submission = &schema.SolutionItem{
		ProblemID: id,
		UserID:    userId,
		Status:    resp.Status,
		Time:      resp.Time,
		Memory:    resp.Memory,
		Indate:    time.Now(),
		Language:  "C",
		Passrate:  resp.RunTime, // 这里为什么使用 RunTime ？
	}
	//改成solution的creat创建
	err = SolutionSvc.Update(submissionId, submission)
	if err != nil {
		logs.Error("Failed to update submission:", err)
		return 0, err
	}

	// Step 7: 返回提交记录的 ID
	return submissionId, nil
}
