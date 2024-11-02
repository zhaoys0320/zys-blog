package views

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"zys-boke-master/common"
	"zys-boke-master/config"
	"zys-boke-master/dao"
	"zys-boke-master/models"
	"zys-boke-master/service"
)

var HTML = &HTMLApi{}

type HTMLApi struct {
}

var API = &Api{}

type Api struct {
}

func PostDetail(w http.ResponseWriter, r *http.Request) {
	//if err := r.ParseForm(); err != nil{
	//	common.Error(w,errors.New("参数解析错误"))
	//	log.Println(err)
	//	return
	//}
	path := r.URL.Path
	id := strings.TrimPrefix(path, "/p/")
	id = strings.TrimSuffix(id, ".html")
	pid, _ := strconv.Atoi(id)
	post, err := dao.GetPostById(pid)
	if err != nil {
		log.Println(err)
		return
	}
	var pm models.PostMore
	pm.UserName = dao.GetUserNameById(post.UserId)
	pm.Pid = post.Pid
	pm.ViewCount = post.ViewCount
	pm.CategoryId = post.CategoryId
	pm.CategoryName = dao.GetCategoryNameById(post.CategoryId)
	pm.Content = template.HTML(post.Content)
	pm.Title = post.Title
	pm.Slug = post.Slug
	pm.CreateAt = common.Format(post.CreateAt)
	common.Template.Detail.WriteData(w,
		models.PostRes{
			config.Cfg.Viewer,
			config.Cfg.System,
			pm,
		})
}

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	id := strings.TrimPrefix(path, "/c/")
	cId, _ := strconv.Atoi(id)
	_ = r.ParseForm()
	page := r.Form.Get("page")
	if page == "" {
		page = "1"
	}
	currentPage, _ := strconv.Atoi(page)
	cName := dao.GetCategoryNameById(cId)
	categorys := dao.GetCategorys()
	post, total := service.PostPageByCategory(currentPage, 10, cId)
	pagesAll := ((total - 1) / 10) + 1
	pages := []int{}
	for i := 1; i <= pagesAll; i++ {
		pages = append(pages, i)
	}
	hd := models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		post,
		total,
		currentPage,
		pages,
		currentPage != pagesAll,
	}
	var categoryData = &models.CategoryData{
		hd,
		cName,
	}

	common.Template.Category.WriteData(w, categoryData)
}
