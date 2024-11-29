package schema

import (
	"github.com/jinzhu/copier"
)

type UsersItem struct {
	ID         int `json:"id"`
	Name       string
	Roles      string
	Email      string
	Submit     int
	Solved     int
	Password   string
	School     string
	Accesstime string
}

type UsersDBItem struct {
	ID         int
	Name       string
	Roles      string
	Email      string
	Submit     int
	Solved     int
	Password   string
	School     string
	Accesstime string
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

func (a *UsersDBItem) TableName() string {
	return "users"
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
	Name string
	School string
	Page int
	PageSize int
}
