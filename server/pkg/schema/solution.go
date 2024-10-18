package schema

import (
	"github.com/jinzhu/copier"
)

type SolutionItem struct {
	ID int
	Problem_ID  int
	User_ID     string
	Time        int
	Memory      int
	In_date     string
	Language    string
	Code_length string
	Juage_time  string
	Juager      string
	Pass_rate   string
}

type SolutionDBItem struct {
	ID int
	Problem_ID  int
	User_ID     string
	Time        int
	Memory      int
	In_date     string
	Language    string
	Code_length string
	Juage_time  string
	Juager      string
	Pass_rate   string
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
}
