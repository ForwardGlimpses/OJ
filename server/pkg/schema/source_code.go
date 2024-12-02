package schema

import (
	"github.com/jinzhu/copier"
)

type SourceCodeItem struct {
	ID     int `json:"id"`
	Source string
}

type SourceCodeDBItem struct {
	ID     int
	Source string
}

func (a *SourceCodeItem) ToDBItem() *SourceCodeDBItem {
	ret := &SourceCodeDBItem{}
	copier.Copy(ret, a)
	return ret
}

func (a *SourceCodeDBItem) TableName() string {
	return "source_code"
}

type SourceCodeItems []*SourceCodeItem

func (a SourceCodeItems) ToDBItems() SourceCodeDBItems {
	ret := SourceCodeDBItems{}
	for _, t := range a {
		ret = append(ret, t.ToDBItem())
	}
	return ret
}

func (a *SourceCodeDBItem) ToItem() *SourceCodeItem {
	ret := &SourceCodeItem{}
	copier.Copy(ret, a)
	return ret
}

type SourceCodeDBItems []*SourceCodeDBItem

func (a SourceCodeDBItems) ToItems() SourceCodeItems {
	ret := SourceCodeItems{}
	for _, t := range a {
		ret = append(ret, t.ToItem())
	}
	return ret
}

type SourceCodeParams struct {
	SolutionID int
	P
}
