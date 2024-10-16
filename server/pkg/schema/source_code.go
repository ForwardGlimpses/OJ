package schema

import (
	"github.com/jinzhu/copier"
)

type Source_codeItem struct {
	ID int
	Source   string
}

type Source_codeDBItem struct {
    ID int
	Source      string
}

func (a *Source_codeItem) ToDBItem() *Source_codeItem {
	ret := &Source_codeItem{}
	copier.Copy(ret, a)
	return ret
}

type SourceCodeItems []*Source_codeItem

func (a SourceCodeItems) ToDBItems() Source_codeItems {
	ret := Source_codeItems{}
	for _, t := range a {
		ret = append(ret, t.ToDBItem())
	}
	return ret
}

func (a *Source_codeItem) ToItem() *Source_codeItem {
	ret := &Source_codeItem{}
	copier.Copy(ret, a)
	return ret
}

type Source_codeItems []*Source_codeItem

func (a Source_codeItems) ToItems() SourceCodeItems {
	ret := SourceCodeItems{}
	for _, t := range a {
		ret = append(ret, t.ToItem())
	}
	return ret
}

type SourceCodeParams struct {
	SolutionID int
}
