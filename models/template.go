package models

import (
	"html/template"
	"net/http"
	"time"
	"zys-boke-master/config"
)

type TemplateBlog struct {
	*template.Template
}
type HtmlTemplate struct {
	Index      TemplateBlog
	Category   TemplateBlog
	Custom     TemplateBlog
	Detail     TemplateBlog
	Login      TemplateBlog
	PigeOnhole TemplateBlog
	Writting   TemplateBlog
}

func (t *TemplateBlog) WriteData(w http.ResponseWriter, data interface{}) {
	t.Execute(w, data)
}
func Date(string2 string) string {
	return time.Now().Format(string2)
}

func GetNextName(strs []string, a int) string {
	return strs[a+1]
}

func IsODD(num int) bool {
	return num%2 == 0
}
func CreateAt(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func InitTemplate() (HtmlTemplate, error) {
	templatehtml := HtmlTemplate{}
	tempBlogs, err := ReadTemplate(
		config.Cfg.System.CurrentDir+"/template/",
		[]string{"index", "category", "custom", "detail", "login", "pigeOnhole", "writing"})
	if err != nil {
		return templatehtml, err
	}
	templatehtml.Index = tempBlogs[0]
	templatehtml.Category = tempBlogs[1]
	templatehtml.Custom = tempBlogs[2]
	templatehtml.Detail = tempBlogs[3]
	templatehtml.Login = tempBlogs[4]
	templatehtml.PigeOnhole = tempBlogs[5]
	templatehtml.Writting = tempBlogs[6]
	return templatehtml, nil
}

func ReadTemplate(templateDir string, templateName []string) ([]TemplateBlog, error) {
	templateblogs := []TemplateBlog{}

	for _, name := range templateName {
		tempName := templateDir + name + ".html"
		t := template.New(name + ".html")
		home := templateDir + "home.html"
		header := templateDir + "Layout/header.html"
		footer := templateDir + "layout/footer.html"
		personal := templateDir + "layout/personal.html"
		post := templateDir + "layout/post-list.html"
		pagination := templateDir + "layout/pagination.html"
		t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date, "dateDay": CreateAt})
		t, err := t.ParseFiles(tempName, home, header, footer, personal, post, pagination)
		if err != nil {
			return templateblogs, err
		}
		templateblogs = append(templateblogs, TemplateBlog{t})
	}
	return templateblogs, nil
}
