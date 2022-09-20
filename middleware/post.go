package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func WithPostData(w http.ResponseWriter, r *http.Request, data interface{}) error {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return errors.New(fmt.Sprintf("method %s not allowed", r.Method))
	}
	body, _ := ioutil.ReadAll(r.Body)
	//var data api.JsonRequest
	if err := json.Unmarshal(body, &data); err != nil {
		return errors.New(fmt.Sprintf("unmarshal error: %v", err))
	}
	return nil
}
