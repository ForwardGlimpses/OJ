package schema

import (
	"github.com/jinzhu/copier"
)

type ContestParams struct {
	Title string
}

type ContestItem struct {
	ID  int
	Title       string
	Start_time  string
	End_time    string
	Password    string
	Administrator string
	Description string
}

type ContestDBItem struct {
	ID  int
	Title       string
	Start_time  string
	End_time    string
	Password    string
	Administrator string
	Description string
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