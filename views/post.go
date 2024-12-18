package views

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"zys-boke-master/common"
	"zys-boke-master/dao"
	"zys-boke-master/models"
)

func AddOrUpdate(w http.ResponseWriter, r *http.Request) {
	//先判断是否为POST还是PUT POST为新增 PUT为更新
	method := r.Method
	switch method {
	case http.MethodPost:
		//获取登录用户信息
		token := r.Header.Get("Authorization")
		_, claims, err := common.ParseToken(token)
		if err != nil {
			log.Println(err)
			return
		}
		uid := claims.Uid
		//获取参数
		param := common.GetRequestJsonParam(r)
		categoryId := param["categoryId"].(string)
		content := param["content"].(string)
		markdown := param["markdown"].(string)
		slug := param["slug"].(string)
		title := param["title"].(string)
		articleType := float64(0)
		if param["type"] != nil {
			articleType = param["type"].(float64)
		}
		post := new(models.Post)
		post.Title = title
		post.UserId = uid
		post.ViewCount = 0
		cId, _ := strconv.Atoi(categoryId)
		post.CategoryId = cId
		post.Markdown = markdown
		post.Slug = slug
		post.Type = int(articleType)
		post.Content = content
		post.CreateAt = time.Now()
		post.UpdateAt = time.Now()
		if err := dao.SavePost(post); err != nil {
			log.Println(err)
			return
		}
		common.Success(w, post)
		return
	case http.MethodPut:
		//获取登录用户信息
		token := r.Header.Get("Authorization")
		_, _, err := common.ParseToken(token)
		if err != nil {
			log.Println(err)
			return
		}
		//获取参数
		param := common.GetRequestJsonParam(r)
		userId := param["userId"].(float64)
		categoryId := param["categoryId"].(float64)
		content := param["content"].(string)
		markdown := param["markdown"].(string)
		slug := param["slug"].(string)
		title := param["title"].(string)
		articleType := float64(0)
		if param["type"] != nil {
			articleType = param["type"].(float64)
		}
		pid := param["pid"].(float64)
		post := new(models.Post)
		post.Pid = int(pid)
		post.Title = title
		post.UserId = int(userId)
		post.CategoryId = int(categoryId)
		post.Markdown = markdown
		post.Slug = slug
		post.Type = int(articleType)
		post.Content = content
		post.CreateAt = time.Now()
		post.UpdateAt = time.Now()
		if err := dao.UpdatePost(post); err != nil {
			log.Println(err)
			return
		}
		common.Success(w, post)
		return

	}
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	//if err := r.ParseForm(); err != nil{
	//	common.Error(w,errors.New("参数解析错误"))
	//	log.Println(err)
	//	return
	//}
	path := r.URL.Path
	id := strings.TrimPrefix(path, "/api/v1/post/")
	pid, _ := strconv.Atoi(id)
	post, err := dao.GetPostById(pid)
	if err != nil {
		log.Println(err)
		return
	}
	common.Success(w, post)
}

func (*Api) PostSearch(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	searchStr := r.Form.Get("val")
	posts := dao.PostSearch(searchStr)
	var searchResp []models.SearchResp
	for _, post := range posts {
		var sr models.SearchResp
		sr.Pid = post.Pid
		sr.Title = post.Title
		searchResp = append(searchResp, sr)
	}
	common.Success(w, searchResp)
}
