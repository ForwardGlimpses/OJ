package service

import (
	"context"
	"sort"
	"time"

	"github.com/ForwardGlimpses/OJ/server/pkg/global"
	"github.com/ForwardGlimpses/OJ/server/pkg/judge"
	"github.com/ForwardGlimpses/OJ/server/pkg/logs"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
)

type ContestSolutionServiceInterface interface {
	Create(ctx context.Context, item *schema.ContestSolutionItem) (int, error)
	Get(ctx context.Context, id int) (*schema.ContestSolutionItem, error)
	Update(ctx context.Context, id int, item *schema.ContestSolutionItem) error
	Delete(ctx context.Context, id int) error
	GetContestSolutions(ctx context.Context, contestID int) (schema.ContestSolutionItems, error)
	GetContestRanking(ctx context.Context, contestID int) ([]ContestRankingItem, error)
	Submit(ctx context.Context, contestID int, input *schema.Submit) (int, error)
}

var ContestSolutionSvc ContestSolutionServiceInterface = &ContestSolutionService{}

type ContestRankingItem struct {
	UserID       int
	TotalSolved  int
	TotalPenalty int
	Problems     map[int]ProblemStatus
}

type ProblemStatus struct {
	IsAccepted  bool
	SubmitTime  time.Time
	Attempts    int
	PenaltyTime int
	Status      string
	RunTime     int
	Memory      int
}

type ContestSolutionService struct{}

// Create 创建新的比赛解决方案
func (a *ContestSolutionService) Create(ctx context.Context, item *schema.ContestSolutionItem) (int, error) {
	db := global.DB.WithContext(ctx)
	dbItem := item.ToDBItem()
	err := db.Create(dbItem).Error
	if err != nil {
		logs.Error("Failed to create contest solution:", err)
		return 0, err
	}
	return dbItem.ID, nil
}

// Get 获取比赛解决方案信息
func (a *ContestSolutionService) Get(ctx context.Context, id int) (*schema.ContestSolutionItem, error) {
	db := global.DB.WithContext(ctx)
	item := &schema.ContestSolutionDBItem{}
	err := db.Where("id = ?", id).First(item).Error
	if err != nil {
		logs.Error("Failed to get contest solution with ID:", id, "Error:", err)
		return nil, err
	}
	return item.ToItem(), nil
}

// Update 更新比赛解决方案信息
func (a *ContestSolutionService) Update(ctx context.Context, id int, item *schema.ContestSolutionItem) error {
	db := global.DB.WithContext(ctx)
	dbItem := item.ToDBItem()
	err := db.Where("id = ?", id).Updates(dbItem).Error
	if err != nil {
		logs.Error("Failed to update contest solution with ID:", id, "Error:", err)
		return err
	}
	return nil
}

// Delete 删除比赛解决方案
func (a *ContestSolutionService) Delete(ctx context.Context, id int) error {
	db := global.DB.WithContext(ctx)
	err := db.Where("id = ?", id).Delete(&schema.ContestSolutionDBItem{}).Error
	if err != nil {
		logs.Error("Failed to delete contest solution with ID:", id, "Error:", err)
		return err
	}
	return nil
}

// GetContestSolutions 获取比赛的所有解决方案
func (a *ContestSolutionService) GetContestSolutions(ctx context.Context, contestID int) (schema.ContestSolutionItems, error) {
	var solutions schema.ContestSolutionDBItems
	err := global.DB.WithContext(ctx).Where("contest_id = ?", contestID).Order("submit_time DESC").Find(&solutions).Error
	if err != nil {
		logs.Error("Failed to get contest solutions:", err)
		return nil, err
	}
	return solutions.ToItems(), nil
}

// GetContestRanking 获取比赛的实时排名
func (a *ContestSolutionService) GetContestRanking(ctx context.Context, contestID int) ([]ContestRankingItem, error) {
	// Step 1: 获取比赛的所有解决方案
	solutions, err := a.GetContestSolutions(ctx, contestID)
	if err != nil {
		return nil, err
	}

	// Step 2: 初始化排名数据结构
	ranking := make(map[int]*ContestRankingItem)

	// Step 3: 遍历所有解决方案，计算每个用户的排名信息
	for _, solution := range solutions {
		// 如果用户还没有在排名中，初始化用户的排名信息
		if _, exists := ranking[solution.UserID]; !exists {
			ranking[solution.UserID] = &ContestRankingItem{
				UserID:       solution.UserID,
				TotalSolved:  0,
				TotalPenalty: 0,
				Problems:     make(map[int]ProblemStatus),
			}
		}

		// 获取用户的排名信息
		userRanking := ranking[solution.UserID]
		// 获取用户在该问题上的状态
		problemStatus := userRanking.Problems[solution.ProblemID]

		var contest schema.ContestDBItem
		err := global.DB.WithContext(ctx).Where("id = ?", contestID).First(&contest).Error
		if err != nil {
			logs.Error("Failed to get contest start time:", err)
			return nil, err
		}

		// 如果解决方案被接受
		if solution.IsAccepted {
			// 如果该问题还没有被接受过
			if !problemStatus.IsAccepted {
				// 增加用户解决的问题数量
				userRanking.TotalSolved++
				// 增加用户的总罚时
				// 罚时 = 提交错误的罚时 + 从比赛开始到解决成功所用时间
				userRanking.TotalPenalty += problemStatus.PenaltyTime*(problemStatus.Attempts-1) + int(solution.SubmitTime.Sub(contest.StartTime).Minutes())
				// 更新问题状态为已接受
				problemStatus.IsAccepted = true
				// 记录提交时间
				problemStatus.SubmitTime = solution.SubmitTime
				problemStatus.PenaltyTime = solution.PenaltyTime
				problemStatus.Status = solution.Status
				problemStatus.RunTime = solution.RunTime
				problemStatus.Memory = solution.Memory
			}
		} else {
			// 如果解决方案没有被接受，增加尝试次数
			problemStatus.Attempts++
		}

		// 更新用户在该问题上的状态
		userRanking.Problems[solution.ProblemID] = problemStatus
	}

	// Step 4: 将排名数据转换为列表
	var rankingList []ContestRankingItem
	for _, item := range ranking {
		rankingList = append(rankingList, *item)
	}

	// 排序排名列表
	sort.Slice(rankingList, func(i, j int) bool {
		if rankingList[i].TotalSolved != rankingList[j].TotalSolved {
			return rankingList[i].TotalSolved > rankingList[j].TotalSolved
		}
		return rankingList[i].TotalPenalty < rankingList[j].TotalPenalty
	})

	// Step 5: 返回排名列表
	return rankingList, nil
}

// Submit 提交解决方案
func (a *ContestSolutionService) Submit(ctx context.Context, contestID int, input *schema.Submit) (int, error) {
	// Step 1: 获取题目详细信息
	problem, err := ProblemSvc.Get(ctx, input.ID)
	if err != nil {
		logs.Error("Failed to get problem with ID:", input.ID, "Error:", err)
		return 0, err
	}

	submission := &schema.ContestSolutionItem{
		ContestID:  contestID,
		ProblemID:  contestID,
		UserID:     input.ID,
		SubmitTime: time.Now(),
		Status:     "Pending",
	}

	submissionId, err := ContestSolutionSvc.Create(ctx, submission)
	if err != nil {
		logs.Error("Failed to create submission:", err)
		return 0, err
	}

	resp, err := judge.Submit(judge.Request{
		ID:     submissionId,
		Code:   input.InputCode,
		Input:  problem.Input,
		Output: problem.Output,
	})

	if err != nil {
		logs.Error("Failed to judge:", err)
		return 0, err
	}

	// Step 6: 将评测结果存储到数据库
	submission = &schema.ContestSolutionItem{
		ContestID: contestID,
		ProblemID: contestID,
		UserID:    input.ID,
		Status:    resp.Status,
	}
	err = ContestSolutionSvc.Update(ctx, submissionId, submission)
	if err != nil {
		logs.Error("Failed to update submission:", err)
		return 0, err
	}

	return submissionId, nil
}
