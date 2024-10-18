package schema

import (
	"github.com/jinzhu/copier"
)

type UsersItem struct {
	ID    int
	Email       string
	Sumbit      int
	Solve       int
	Password    string
	School      string
	Access_time string
	Enroll_time string
}

type UsersDBItem struct {
	ID    int
	Email       string
	Sumbit      int
	Solve       int
	Password    string
	School      string
	Access_time string
	Enroll_time string
}

func (a *UsersItem) ToDBItem() *UsersDBItem {
	ret := &UsersDBItem{}
	copier.Copy(ret, a)
	return ret
}

type UsersItems []*UsersItem

func (a UsersItems) ToDBItems() UsersDBItems {
	ret := UsersDBItems{}
	for _, t := range a {
		ret = append(ret, t.ToDBItem())
	}
	return ret
}

func (a *UsersDBItem) ToItem() *UsersItem {
	ret := &UsersItem{}
	copier.Copy(ret, a)
	return ret
}

type UsersDBItems []*UsersDBItem

func (a UsersDBItems) ToItems() UsersItems {
	ret := UsersItems{}
	for _, t := range a {
		ret = append(ret, t.ToItem())
	}
	return ret
}

type UsersParams struct {
	Email string
}
