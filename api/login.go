package api

import (
	"net/http"
	"zys-boke-master/common"
	"zys-boke-master/service"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {

	params := common.GetRequestJsonParam(r)
	un := params["username"].(string)
	wd := params["passwd"].(string)
	userRes := service.Login(un, wd)
	common.Success(w, userRes)
}
