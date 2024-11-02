package views

import (
	"net/http"
	"zys-boke-master/common"
	"zys-boke-master/config"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	login := common.Template.Login

	login.Execute(w, config.Cfg.Viewer)
}
