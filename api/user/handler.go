package user

import (
	"context"
	"github.com/MelkoV/go-learn-api/rpc"
	_ "github.com/MelkoV/go-learn-api/rpc"
	"github.com/MelkoV/go-learn-logger/logger"
)

type Login struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	IsRemember bool   `json:"isRemember"`
}

func (lr *Login) Handler(ctx context.Context, l *logger.CategoryLogger) {
	l.Format("user", ctx.Value(rpc.CtxUuidKey).(string), "test: %v", true).Info()
}
