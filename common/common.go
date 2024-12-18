package common

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
	"zys-boke-master/models"
)

var Template models.HtmlTemplate

func Md5Crypt(str string, salt ...interface{}) (CryptStr string) {
	if l := len(salt); l > 0 {
		slice := make([]string, l+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
func LoadTemplate() {
	wg := sync.WaitGroup{}
	//加载html模板
	var err error
	wg.Add(1)
	go func() {
		Template, err = models.InitTemplate()
		if err != nil {
			panic(err)
		}
		wg.Done()
	}()
	wg.Wait()
}
func Success(w http.ResponseWriter, data interface{}) {
	var ret models.Result
	ret.Code = 200
	ret.Data = data
	ret.Error = ""
	resultJson, _ := json.Marshal(ret)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(resultJson)
	if err != nil {
		log.Fatal(err)
	}
}
func Format(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}

func FormatMonth(time time.Time) string {
	return time.Format("2006-01")
}
func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &params)
	return params
}
