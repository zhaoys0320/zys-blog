package models

type Category struct {
	Cid      int
	Name     string
	CreateAt string
	UpdateAt string
}

type CategoryData struct {
	HomeResponse
	CategoryName string
}
