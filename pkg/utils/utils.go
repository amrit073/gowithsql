package utils

import (
	"encoding/json"
	"fmt"
	//"fmt"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func First[T any](a T, e T) T {

	fmt.Println(e)
	return a
}
