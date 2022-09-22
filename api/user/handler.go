package user

import (
	"context"
	"github.com/MelkoV/go-learn-api/api/service"
	"github.com/MelkoV/go-learn-api/rpc"
	"github.com/MelkoV/go-learn-logger/logger"
	pb "github.com/MelkoV/go-learn-proto/proto/user"
	"net/http"
)

func (m *LoginRequest) Handle(ctx context.Context, l logger.CategoryLogger, w http.ResponseWriter, r *http.Request) {
	l.Debug("model %v", m)
	return
	uuid := ctx.Value(rpc.CtxUuidKey).(string)
	client := service.NewUserServiceClient(ctx, l)
	user, err := client.Login(ctx, &pb.LoginRequest{
		Uuid: uuid,
		User: &pb.User{
			Username: m.Login.Username,
			Password: m.Login.Password,
		},
		Remember: m.Login.IsRemember,
	})
	if err != nil {
		l.Error("error call user service: %v", err)
	}
	l.Info("user result: %v", user)
}

type Status struct {
	Session string `json:"session"`
}

func (m *Status) Handler(ctx context.Context, l logger.CategoryLogger, w http.ResponseWriter, r *http.Request) {
	//uuid := ctx.Value(rpc.CtxUuidKey).(string)

}
