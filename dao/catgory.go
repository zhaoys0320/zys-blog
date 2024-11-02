package dao

import (
	"log"
	"zys-boke-master/models"
)

func GetCategorys() []models.Category {
	ret, err := DB.Query("select * from blog_category")
	if err != nil {
		log.Println(err)
		return nil
	}
	var cs []models.Category
	for ret.Next() {
		var cat models.Category
		err = ret.Scan(&cat.Name, &cat.CreateAt, &cat.UpdateAt, &cat.Cid)
		if err != nil {
			log.Println(err)
		}
		cs = append(cs, cat)
	}
	return cs
}
func GetCategoryNameById(id int) string {
	row := DB.QueryRow("select name from blog_category where id=?", id)
	var name string
	_ = row.Scan(&name)
	return name
}
