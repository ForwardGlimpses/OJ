package schema

import (
	"time"

	"github.com/jinzhu/copier"
)

type ContestItem struct {
	ID            int `json:"id"`
	Title         string
	Private       string
	StartTime     time.Time
	EndTime       time.Time
	Password      string
	Administrator string
	Description   string
}

type ContestDBItem struct {
	ID            int
	Title         string
	Private       string
	Start_time    time.Time
	End_time      time.Time
	Password      string
	Administrator string
	Description   string
}

func (a *ContestItem) ToDBItem() *ContestDBItem {
	ret := &ContestDBItem{}
	copier.Copy(ret, a)
	return ret
}

type ContestItems []*ContestItem

func (a ContestItems) ToDBItems() ContestDBItems {
	ret := ContestDBItems{}
	for _, t := range a {
		ret = append(ret, t.ToDBItem())
	}
	return ret
}

func (a *ContestDBItem) TableName() string {
	return "contest"
}

func (a *ContestDBItem) ToItem() *ContestItem {
	ret := &ContestItem{}
	copier.Copy(ret, a)
	return ret
}

type ContestDBItems []*ContestDBItem

func (a ContestDBItems) ToItems() ContestItems {
	ret := ContestItems{}
	for _, t := range a {
		ret = append(ret, t.ToItem())
	}
	return ret
}

type ContestParams struct {
	Title string
	P
}
