package schema

import (
	"github.com/jinzhu/copier"
)

type PaginatedResultItem struct {
    Items      []ContestItem  // 当前页的数据
    TotalCount int64          // 总记录数
    Page       int            // 当前页码
    PageSize   int            // 每页条数
}

func (a *PaginatedResultItem) ToDBItem() *PaginatedResultDBItem {
	ret := &PaginatedResultDBItem{}
	copier.Copy(ret, a)
	return ret
}

type PaginatedResultItems []*PaginatedResultItem

func (a PaginatedResultItems) ToDBItems() PaginatedResultDBItems {
	ret := PaginatedResultDBItems{}
	for _, t := range a {
		ret = append(ret, t.ToDBItem())
	}
	return ret
}

type PaginatedResultDBItem struct {
	Items      []ContestItem  // 当前页的数据
    TotalCount int64          // 总记录数
    Page       int            // 当前页码
    PageSize   int            // 每页条数
}

func (a *PaginatedResultDBItem) TableName() string {
	return "paginatedresult"
}

func (a *PaginatedResultDBItem) ToItem() *PaginatedResultItem {
	ret := &PaginatedResultItem{}
	copier.Copy(ret, a)
	return ret
}

type PaginatedResultDBItems []*PaginatedResultDBItem

func (a PaginatedResultDBItems) ToItems() PaginatedResultItems {
	ret := PaginatedResultItems{}
	for _, t := range a {
		ret = append(ret, t.ToItem())
	}
	return ret
}

type PaginatedResultParams struct {
	Title string
}
