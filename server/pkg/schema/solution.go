package schema

import (
	"github.com/jinzhu/copier"
)

type SolutionItem struct {
	ID         int `json:"id"`
	ProblemID  int
	UserID     int
	Time       uint64
	Memory     uint64
	Status     string
	InDate     string
	Language   string
	Codelength string
	Judgetime  string
	Judger     string
	Passrate   string
}

type SolutionDBItem struct {
	ID         int `gorm:"primary_key;auto_increment"`
	ProblemID  int
	UserID     int
	Time       uint64
	Memory     uint64
	Status     string
	Indate     string
	Language   string
	Codelength string
	Judgetime  string
	Judger     string
	Passrate   string
}

func (a *SolutionItem) ToDBItem() *SolutionDBItem {
	ret := &SolutionDBItem{}
	copier.Copy(ret, a)
	return ret
}

type SolutionItems []*SolutionItem

func (a SolutionItems) ToDBItems() SolutionDBItems {
	ret := SolutionDBItems{}
	for _, t := range a {
		ret = append(ret, t.ToDBItem())
	}
	return ret
}

func (a *SolutionDBItem) TableName() string {
	return "solution"
}

func (a *SolutionDBItem) ToItem() *SolutionItem {
	ret := &SolutionItem{}
	copier.Copy(ret, a)
	return ret
}

type SolutionDBItems []*SolutionDBItem

func (a SolutionDBItems) ToItems() SolutionItems {
	ret := SolutionItems{}
	for _, t := range a {
		ret = append(ret, t.ToItem())
	}
	return ret
}

type SolutionParams struct {
	UserID string
	P
}
