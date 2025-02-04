package service

import (
	"context"
	"fmt"
	"time"

	"github.com/ForwardGlimpses/OJ/server/pkg/global"
	"github.com/ForwardGlimpses/OJ/server/pkg/gormx"
	"github.com/ForwardGlimpses/OJ/server/pkg/judge"
	"github.com/ForwardGlimpses/OJ/server/pkg/logs"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
)

type ProblemServiceInterface interface {
	Query(ctx context.Context, params schema.ProblemParams) (schema.ProblemItems, int64, error)
	Get(ctx context.Context, id int) (*schema.ProblemItem, error)
	Create(ctx context.Context, item *schema.ProblemItem) (int, error)
	Update(ctx context.Context, id int, item *schema.ProblemItem) error
	Delete(ctx context.Context, id int) error
	Submit(ctx context.Context, id int, userId int, inputCode string) (int, error)
}

var ProblemSvc ProblemServiceInterface = &ProblemService{}

type ProblemService struct{}

// Query 根据条件和分页查询获取题目列表
func (a *ProblemService) Query(ctx context.Context, params schema.ProblemParams) (schema.ProblemItems, int64, error) {
	// 初始化查询
	query := global.DB.WithContext(ctx).Model(&schema.ProblemDBItem{})

	// 应用过滤条件
	if params.ProblemID != 0 {
		query = query.Where("problem_id = ?", params.ProblemID)
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
func (a *ProblemService) Get(ctx context.Context, id int) (*schema.ProblemItem, error) {
	db := global.DB.WithContext(ctx)
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
func (a *ProblemService) Create(ctx context.Context, item *schema.ProblemItem) (int, error) {
	db := global.DB.WithContext(ctx)
	dbItem := item.ToDBItem()
	err := db.Create(dbItem).Error
	if err != nil {
		logs.Error("Failed to create problem:", err)
		return 0, err
	}
	return dbItem.ID, nil
}

// Update 更新题目信息
func (a *ProblemService) Update(ctx context.Context, id int, item *schema.ProblemItem) error {
	db := global.DB.WithContext(ctx)
	dbItem := item.ToDBItem()
	err := db.Where("id = ?", id).Updates(dbItem).Error
	if err != nil {
		logs.Error("Failed to update problem with ID:", id, "Error:", err)
		return err
	}
	return nil
}

// Delete 根据ID删除题目
func (a *ProblemService) Delete(ctx context.Context, id int) error {
	db := global.DB.WithContext(ctx)
	err := db.Where("id = ?", id).Delete(&schema.ProblemDBItem{}).Error
	if err != nil {
		logs.Error("Failed to delete problem with ID:", id, "Error:", err)
		return err
	}
	return nil
}

// Submit 提交解决方案
func (a *ProblemService) Submit(ctx context.Context, id int, userId int, code string) (int, error) {
	// Step 1: 获取题目详细信息
	problem, err := a.Get(ctx, id)
	if err != nil {
		logs.Error("Failed to get problem with ID:", id, "Error:", err)
		return 0, err
	}

	submission := &schema.SolutionItem{
		ProblemID: id,
		UserID:    userId,
		Status:    "Pending",
	}

	submissionId, err := SolutionSvc.Create(ctx, submission)
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
		Passrate:  resp.RunTime,
	}
	err = SolutionSvc.Update(ctx, submissionId, submission)
	if err != nil {
		logs.Error("Failed to update submission:", err)
		return 0, err
	}

	fmt.Println("Submission created with ID:", submissionId)

	// Step 7: 返回提交记录的 ID
	return submissionId, nil
}
