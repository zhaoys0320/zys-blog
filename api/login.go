package api

import (
	"errors"
	"log"
	"net/http"
	"zys-boke-master/common"
	"zys-boke-master/dao"
	"zys-boke-master/models"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	param := common.GetRequestJsonParam(r)
	username := param["username"].(string)
	passwd := param["passwd"].(string)
	passwd = common.Md5Crypt(passwd, "mszlu")
	loginReq := new(models.LoginReq)
	loginReq.Name = username
	loginReq.Passwd = passwd
	user, dbError := dao.Login(loginReq)
	if dbError != nil {
		if dbError.IsNilError {
			dbError.Err = errors.New("账号密码不正确")
		}
		log.Println(dbError)
		return
	}
	uid := user.Uid
	token, _ := common.Award(&uid)
	loginRes := &models.LoginResp{Token: token, UserInfo: models.UserRes{user.Uid, user.UserName, user.Avatar}}
	common.Success(w, loginRes)
}
