package main

import (
	"log"
	"net/http"
	"zys-boke-master/common"
	"zys-boke-master/router"
)

func init() {
	common.LoadTemplate()
}
func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
