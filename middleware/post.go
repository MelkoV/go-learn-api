package middleware

import (
	"errors"
	"fmt"
	"net/http"
)

func WithPostData(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return errors.New(fmt.Sprintf("method %s not allowed", r.Method))
	}
	
	return nil
}
