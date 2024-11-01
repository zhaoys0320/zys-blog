package router

import (
	"net/http"
	"zys-boke-master/views"
)

func Router() {

	http.HandleFunc("/", views.Html.MyHandle)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
