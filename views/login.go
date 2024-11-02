package views

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"strings"
	"zys-boke-master/common"
	"zys-boke-master/config"
)

func LoginHtml(w http.ResponseWriter, r *http.Request) {
	login := common.Template.Login
	login.WriteData(w, config.Cfg.Viewer)
}

func Md5Crypt(str string, salt ...interface{}) (CryptStr string) {
	if l := len(salt); l > 0 {
		slice := make([]string, l+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	//param := common.GetRequestJsonParam(r)
	//username := param["username"].(string)
	//passwd := param["passwd"].(string)
	//passwd = Md5Crypt(passwd, "mszlu")
	//loginReq := new(models.LoginReq)
	//loginReq.Name = username
	//loginReq.Passwd = passwd
	//user, dbError := dao.Login(loginReq)
	//if dbError != nil {
	//	if dbError.IsNilError {
	//		dbError.Err = errors.New("账号密码不正确")
	//	}
	//	return
	//}
	//uid := user.Uid
	//token, _ := common.Award(&uid)
	//loginRes := &models.LoginResp{Token: token, UserInfo: models.UserRes{user.Uid, user.UserName, user.Avatar}}
	//common.Success(w, loginRes)
	login := common.Template.Login
	login.WriteData(w, config.Cfg.Viewer)
}
