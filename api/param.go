package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getGETParam(req *http.Request) map[string]string {
	ret := make(map[string]string)

	req.ParseForm()
	for k, v := range req.Form {
		ret[k] = v[0]
	}
	return ret
}

func getPOSTParam(req *http.Request) (ret map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			ret = make(map[string]interface{})
		}
	}()

	ret = make(map[string]interface{})
	if req == nil {
		return
	}
	if req.Body == nil {
		return
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}

	fmt.Println(string(body))
	err = json.Unmarshal(body, &ret)
	if err != nil {
		fmt.Println(err)
		return make(map[string]interface{})
	}
	return
}
