package main

import (
	"log"
	"net/http"
	"zys-boke-master/common"
	"zys-boke-master/router"
)

type Data struct {
	Title string `json:"title"`
	Des   string `json:"description"`
}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	common.LoadTemplate()
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
