package rpc

import (
	"context"
	"encoding/json"
	"github.com/MelkoV/go-learn-logger/logger"
	"net/http"
)

type ctxKey int

const (
	CtxUuidKey         ctxKey = 1
	MethodHeader              = "X-RPC-METHOD"
	CodeServerError           = 500
	CodeNotFound              = 404
	MessageServerError        = "Server error"
	MessageNotFound           = "Not found"
)

type Action interface {
	Handle(ctx context.Context, l logger.CategoryLogger, w http.ResponseWriter, r *http.Request)
}

type ProtocolError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ProtocolResponse struct {
	Error  ProtocolError `json:"error"`
	Result interface{}   `json:"result"`
}

func WriteError(w http.ResponseWriter, code int, message string) {
	r := ProtocolResponse{
		Error: ProtocolError{
			Code:    code,
			Message: message,
		},
		Result: nil,
	}
	Write(w, r)
}

func WriteOk(w http.ResponseWriter, v interface{}) {
	r := ProtocolResponse{
		Error:  ProtocolError{},
		Result: v,
	}
	Write(w, r)
}

func Write(w http.ResponseWriter, r ProtocolResponse) {
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(r)
	w.Write(data)
}

func SetCookie(w http.ResponseWriter, name string, value string, maxAge int) {
	c := http.Cookie{
		Name:   name,
		Value:  value,
		MaxAge: maxAge,
	}
	http.SetCookie(w, &c)
}
