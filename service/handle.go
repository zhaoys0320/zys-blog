package service

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"zys-boke-master/common"
	"zys-boke-master/config"
	"zys-boke-master/dao"
	"zys-boke-master/models"
)

func Handle(r *http.Request) *models.HomeResponse {
	//path := r.URL.Path
	if err := r.ParseForm(); err != nil {
		log.Println("ParseForm() err:", err)
	}
	page := r.Form.Get("page")
	if page == "" {
		page = "2"
	}
	categorys := dao.GetCategorys()
	currentPage, _ := strconv.Atoi(page)
	//var posts = dao.GetPostAll()
	var postMores = []models.PostMore{}
	posts := dao.GetPostPage(currentPage, 3)
	for _, post := range posts {
		caName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[:100]
		}
		var postMore = models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(content),
			post.CategoryId,
			caName,
			post.UserId,
			userName,
			post.ViewCount,
			post.Type,
			post.CreateAt.String(),
			post.UpdateAt.String(),
		}
		postMores = append(postMores, postMore)
	}
	total := dao.GetPostCount()
	pageSize := 3
	pageCount := ((total - 1) / pageSize) + 1
	var pages []int
	for i := 1; i <= pageCount; i++ {
		pages = append(pages, i)
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		postMores,
		total,
		currentPage,
		pages,
		currentPage == pageCount,
	}
	return hr
}

func PostPageByCategory(page int, pageSize int, categoryId int) ([]models.PostMore, int) {
	posts := dao.GetPostPageCategory(page, pageSize, categoryId)
	total := dao.GetPostCountCategory(categoryId)
	var postMores []models.PostMore
	for _, post := range posts {
		var pm models.PostMore
		pm.Pid = post.Pid
		pm.ViewCount = post.ViewCount
		pm.CategoryId = post.CategoryId
		pm.CategoryName = dao.GetCategoryNameById(post.CategoryId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = []rune(post.Content)[:100]
		}
		pm.Content = template.HTML(content)
		pm.Title = post.Title
		pm.Slug = post.Slug
		pm.CreateAt = post.CreateAt.String()
		pm.UserName = dao.GetUserNameById(post.UserId)
		postMores = append(postMores, pm)
	}
	return postMores, total
}
func Login(username string, password string) models.LoginRes {
	req := models.LoginReq{
		username,
		password,
	}

	userInfo, err := dao.Login(&req)

	if userInfo == nil || err != nil {
		return models.LoginRes{}
	}
	token, _ := common.Award(&userInfo.Uid)
	return models.LoginRes{
		token,
		*userInfo,
	}
}
