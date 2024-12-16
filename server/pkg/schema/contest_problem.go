package schema

import (
	"github.com/jinzhu/copier"
)

type ContestProblemItem struct {
	ID        int
	ProblemID int
	ContestID string
	Title     string
	Num       int
	Accepted  string
	Submited  string
}

func (a *ContestProblemItem) ToDBItem() *ContestProblemDBItem {
	ret := &ContestProblemDBItem{}
	copier.Copy(ret, a)
	return ret
}

type ContestProblemItems []*ContestProblemItem

func (a ContestProblemItems) ToDBItems() ContestProblemDBItems {
	ret := ContestProblemDBItems{}
	for _, t := range a {
		ret = append(ret, t.ToDBItem())
	}
	return ret
}

type ContestProblemDBItem struct {
	ID        int
	ProblemID int
	ContestID string
	Title     string
	Num       int
	Accepted  string
	Submit    string
}

func (a *ContestProblemDBItem) TableName() string {
	return "contest_problem"
}

func (a *ContestProblemDBItem) ToItem() *ContestProblemItem {
	ret := &ContestProblemItem{}
	copier.Copy(ret, a)
	return ret
}

type ContestProblemDBItems []*ContestProblemDBItem

func (a ContestProblemDBItems) ToItems() ContestProblemItems {
	ret := ContestProblemItems{}
	for _, t := range a {
		ret = append(ret, t.ToItem())
	}
	return ret
}

type ContestProblemParams struct {
	Title string
	P
}
