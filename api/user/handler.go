package user

import (
	"context"
	"github.com/MelkoV/go-learn-api/api/service"
	"github.com/MelkoV/go-learn-api/rpc"
	"github.com/MelkoV/go-learn-logger/logger"
	pb "github.com/MelkoV/go-learn-proto/proto/user"
	"net/http"
)

type Login struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	IsRemember bool   `json:"isRemember"`
}

func (m *Login) Handler(ctx context.Context, l *logger.CategoryLogger, w http.ResponseWriter, r *http.Request) {
	uuid := ctx.Value(rpc.CtxUuidKey).(string)
	client := service.NewUserServiceClient(ctx, l)
	user, err := client.Login(ctx, &pb.LoginRequest{
		Uuid: uuid,
		User: &pb.User{
			Username: m.Username,
			Password: m.Password,
		},
		Remember: m.IsRemember,
	})
	if err != nil {
		l.Format("user", uuid, "error call user service: %v", err).Error()
	}
	l.Format("user", uuid, "user result: %v", user).Info()
}

type Status struct {
	Session string `json:"session"`
}

func (m *Status) Handler(ctx context.Context, l *logger.CategoryLogger, w http.ResponseWriter, r *http.Request) {
	//uuid := ctx.Value(rpc.CtxUuidKey).(string)

}
