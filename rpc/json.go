package rpc

import (
	"context"
	"encoding/json"
	"github.com/MelkoV/go-learn-logger/logger"
)

type ctxKey int

const CtxUuidKey ctxKey = 1

type JsonRequest struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Id      interface{} `json:"id"`
	Params  interface{}
}

type Action interface {
	Handler(ctx context.Context, l *logger.CategoryLogger)
}

func FillParams(data JsonRequest, v Action) error {
	raw, err := json.Marshal(data.Params)
	if err != nil {
		return err
	}
	err = json.Unmarshal(raw, v)
	if err != nil {
		return err
	}
	return nil
}

/*
type ActionData interface {
	GetPb()
}

type Action interface {
	Register(func(l *logger.CategoryLogger, ad ActionData))
	Handle(l *logger.CategoryLogger, ad ActionData)
}

type JsonAction struct {
	Method string
	h      func(l *logger.CategoryLogger, ad ActionData)
}

func (ra *JsonAction) Register(h func(l *logger.CategoryLogger, ad ActionData)) {
	ra.h = h
}

func (ra *JsonAction) Handle(l *logger.CategoryLogger, ad ActionData) {
	ra.h(l, ad)
}



func FillParams(data JsonRequest, v ActionData) error {
	raw, err := json.Marshal(data.Params)
	if err != nil {
		return err
	}
	err = json.Unmarshal(raw, v)
	if err != nil {
		return err
	}
	return nil
}

func NewJsonAction(method string) *JsonAction {
	return &JsonAction{Method: method}
}*/
