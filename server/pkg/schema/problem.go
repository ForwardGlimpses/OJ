package schema

import (
	"time"

	"github.com/jinzhu/copier"
)

type ProblemItem struct {
	ID           int `json:"id"`
	Title        string
	Description  string
	Input        string
	Output       string
	SampleInput  string
	SampleOutput string
	Indate       time.Time
	TimeLimit    string
	MemoryLimit  string
	Accepted     int
	Submited     int
	Solved       int
}

type ProblemDBItem struct {
	ID           int
	Title        string
	Description  string
	Input        string
	Output       string
	SampleInput  string
	SampleOutput string
	Indate       time.Time
	TimeLimit    string
	MemoryLimit  string
	Accepted     int
	Submited     int
	Solved       int
}

func (a *ProblemItem) ToDBItem() *ProblemDBItem {
	ret := &ProblemDBItem{}
	copier.Copy(ret, a)
	return ret
}

type ProblemItems []*ProblemItem

func (a ProblemItems) ToDBItems() ProblemDBItems {
	ret := ProblemDBItems{}
	for _, t := range a {
		ret = append(ret, t.ToDBItem())
	}
	return ret
}

func (a *ProblemDBItem) TableName() string {
	return "problem"
}

func (a *ProblemDBItem) ToItem() *ProblemItem {
	ret := &ProblemItem{}
	copier.Copy(ret, a)
	return ret
}

type ProblemDBItems []*ProblemDBItem

func (a ProblemDBItems) ToItems() ProblemItems {
	ret := ProblemItems{}
	for _, t := range a {
		ret = append(ret, t.ToItem())
	}
	return ret
}

type ProblemParams struct {
	ProblemID int
	Title     string
	P
}

type Submit struct {
	ID        int    `json:"id"`        // 题目 ID
	UserID    int    `json:"userid"`    // 用户 ID
	InputCode string `json:"inputcode"` // 用户提交的代码
}
