package service

import (
	"context"
	"sort"
	"time"

	"github.com/ForwardGlimpses/OJ/server/pkg/global"
	"github.com/ForwardGlimpses/OJ/server/pkg/logs"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
)

type ContestSolutionServiceInterface interface {
	GetContestSolutions(contestID int) (schema.ContestSolutionItems, error)
	GetContestRanking(contestID int) ([]ContestRankingItem, error)
	Create(item *schema.ContestSolutionItem) (int, error)
	Update(id int, item *schema.ContestSolutionItem) error
	Delete(id int) error
	Submit(id int, userId int, inputCode string) (int, error)
}

var ContestSolutionSvc ContestSolutionServiceInterface = &ContestSolutionService{}

type ContestSolutionService struct{}

type ContestRankingItem struct {
	UserID       int
	TotalSolved  int
	TotalPenalty int
	Problems     map[int]ProblemStatus
}

type ProblemStatus struct {
	IsAccepted  bool
	SubmitTime  time.Time
	PenaltyTime int
	Status      string
	RunTime     int
	Memory      int
}

// Create 创建新的比赛解决方案
func (a *ContestSolutionService) Create(item *schema.ContestSolutionItem) (int, error) {
	db := global.DB.WithContext(context.Background())
	dbItem := item.ToDBItem()
	err := db.Create(dbItem).Error
	if err != nil {
		logs.Error("Failed to create contest solution:", err)
		return 0, err
	}
	return dbItem.ID, nil
}

// Update 更新比赛解决方案信息
func (a *ContestSolutionService) Update(id int, item *schema.ContestSolutionItem) error {
	db := global.DB.WithContext(context.Background())
	dbItem := item.ToDBItem()
	err := db.Where("id = ?", id).Updates(dbItem).Error
	if err != nil {
		logs.Error("Failed to update contest solution with ID:", id, "Error:", err)
		return err
	}
	return nil
}

// Delete 删除比赛解决方案
func (a *ContestSolutionService) Delete(id int) error {
	db := global.DB.WithContext(context.Background())
	err := db.Where("id = ?", id).Delete(&schema.ContestSolutionDBItem{}).Error
	if err != nil {
		logs.Error("Failed to delete contest solution with ID:", id, "Error:", err)
		return err
	}
	return nil
}

// GetContestSolutions 获取比赛的所有解决方案
func (a *ContestSolutionService) GetContestSolutions(contestID int) (schema.ContestSolutionItems, error) {
	var solutions schema.ContestSolutionDBItems
	err := global.DB.Where("contest_id = ?", contestID).Order("submit_time DESC").Find(&solutions).Error
	if err != nil {
		logs.Error("Failed to get contest solutions:", err)
		return nil, err
	}
	return solutions.ToItems(), nil
}

// GetContestRanking 获取比赛的实时排名
func (a *ContestSolutionService) GetContestRanking(contestID int) ([]ContestRankingItem, error) {
	solutions, err := a.GetContestSolutions(contestID)
	if err != nil {
		return nil, err
	}

	ranking := make(map[int]*ContestRankingItem)
	for _, solution := range solutions {
		if _, exists := ranking[solution.UserID]; !exists {
			ranking[solution.UserID] = &ContestRankingItem{
				UserID:       solution.UserID,
				TotalSolved:  0,
				TotalPenalty: 0,
				Problems:     make(map[int]ProblemStatus),
			}
		}

		userRanking := ranking[solution.UserID]
		problemStatus := userRanking.Problems[solution.ProblemID]

		if solution.IsAccepted {
			if !problemStatus.IsAccepted {
				userRanking.TotalSolved++
				userRanking.TotalPenalty += solution.PenaltyTime
				problemStatus.IsAccepted = true
				problemStatus.SubmitTime = solution.SubmitTime
				problemStatus.PenaltyTime = solution.PenaltyTime
				problemStatus.Status = solution.Status
				problemStatus.RunTime = solution.RunTime
				problemStatus.Memory = solution.Memory
			}
		} else {
			userRanking.TotalPenalty += solution.PenaltyTime
		}

		userRanking.Problems[solution.ProblemID] = problemStatus
	}

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

	return rankingList, nil
}

func (a *ContestSolutionService) Submit(id int, userId int, inputCode string) (int, error) {
	return ProblemSvc.Submit(id, userId, inputCode)
}
