package views

import (
	"net/http"
	"zys-boke-master/common"
	"zys-boke-master/service"
)

func (h *HTMLApi) MyHandle(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	var hr = service.Handle(r)
	index.WriteData(w, hr)
}
