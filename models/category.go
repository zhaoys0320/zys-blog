package models

type Category struct {
	Cid      int    `db:"id"`
	Name     string `db:"name"`
	CreateAt string `db:"create_at"`
	UpdateAt string `db:"update_at"`
}

type CategoryData struct {
	HomeResponse
	CategoryName string
}
