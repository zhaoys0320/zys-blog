package router

import (
	"net/http"
	"zys-boke-master/api"
	"zys-boke-master/views"
)

func Router() {

	http.HandleFunc("/", views.HTML.MyHandle)
	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/login", views.HTML.Login)
	http.HandleFunc("/writing", views.Writing)
	http.HandleFunc("/api/v1/post", views.AddOrUpdate)
	http.HandleFunc("/api/v1/post/", views.GetPost)
	http.HandleFunc("/p/", views.PostDetail)
	http.HandleFunc("/api/v1/post/search", views.API.PostSearch)
	http.HandleFunc("/api/v1/login", api.HTMLApI.Login)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
