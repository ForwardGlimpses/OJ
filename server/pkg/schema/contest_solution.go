package schema

import (
	"time"

	"github.com/jinzhu/copier"
)

type ContestSolutionItem struct {
	ID          int       `json:"id"`
	ContestID   int       `json:"contest_id"`
	UserID      int       `json:"user_id"`
	ProblemID   int       `json:"problem_id"`
	SubmitTime  time.Time `json:"submit_time"`
	IsAccepted  bool      `json:"is_accepted"`
	PenaltyTime int       `json:"penalty_time"` // 罚时，单位为分钟
	Status      string    // 评测状态
	RunTime     int       // 题目运行时间，单位为毫秒
	Memory      int       // 题目运行内存，单位为KB
}

type ContestSolutionDBItem struct {
	ID          int       `gorm:"column:id"`
	ContestID   int       `gorm:"column:contest_id"`
	UserID      int       `gorm:"column:user_id"`
	ProblemID   int       `gorm:"column:problem_id"`
	SubmitTime  time.Time `gorm:"column:submit_time"`
	IsAccepted  bool      `gorm:"column:is_accepted"`
	PenaltyTime int       `gorm:"column:penalty_time"` // 罚时，单位为分钟
	Status      string    // 评测状态
	RunTime     int       // 题目运行时间，单位为毫秒
	Memory      int       // 题目运行内存，单位为KB
}

func (a *ContestSolutionItem) ToDBItem() *ContestSolutionDBItem {
	ret := &ContestSolutionDBItem{}
	copier.Copy(ret, a)
	return ret
}

func (a *ContestSolutionDBItem) ToItem() *ContestSolutionItem {
	ret := &ContestSolutionItem{}
	copier.Copy(ret, a)
	return ret
}

func (a *ContestSolutionDBItem) TableName() string {
	return "contest_solution"
}

type ContestSolutionItems []*ContestSolutionItem

func (a ContestSolutionItems) ToDBItems() ContestSolutionDBItems {
	ret := ContestSolutionDBItems{}
	for _, t := range a {
		ret = append(ret, t.ToDBItem())
	}
	return ret
}

type ContestSolutionDBItems []*ContestSolutionDBItem

func (a ContestSolutionDBItems) ToItems() ContestSolutionItems {
	ret := ContestSolutionItems{}
	for _, t := range a {
		ret = append(ret, t.ToItem())
	}
	return ret
}
