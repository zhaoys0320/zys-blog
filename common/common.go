package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"zys-boke-master/models"
)

var Template models.HtmlTemplate

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

func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &params)
	return params
}
