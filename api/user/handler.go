package user

import (
	"context"
	"github.com/MelkoV/go-learn-api/api/service"
	"github.com/MelkoV/go-learn-api/rpc"
	"github.com/MelkoV/go-learn-common/dictionary"
	du "github.com/MelkoV/go-learn-common/dictionary/source/user"
	"github.com/MelkoV/go-learn-logger/logger"
	pb "github.com/MelkoV/go-learn-proto/proto/user"
	"net/http"
)

func (m *LoginRequest) Handle(ctx context.Context, l logger.CategoryLogger, d dictionary.IStorage, w http.ResponseWriter, r *http.Request) {
	client, err := service.UserServiceClient()
	if err != nil {
		l.Error("failed connect to gRPC user service: %s", err)
		rpc.WriteRpcError(w, rpc.CodeServerError, rpc.MessageServerError)
		return
	}
	/*cc, err := r.Cookie("test")
	if err != nil {
		l.Error("cookie not found")
	} else {
		l.Info("cookie value is %s", cc.Value)
	}*/
	uuid := ctx.Value(rpc.CtxUuidKey).(string)
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
		rpc.WriteRpcError(w, rpc.CodeServerError, rpc.MessageServerError)
		return
	}
	l.Info("user result: %v", user)
	for _, c := range user.Cookie {
		rpc.SetCookie(w, c.Name, c.Value, int(c.MaxAge))
	}
	rpc.WriteError(w, http.StatusBadRequest, map[string]string{"username": d.Get(dictionary.User, du.IncorrectCredentials)})
	//rpc.WriteOk(w, user.User)
}

type Status struct {
	Session string `json:"session"`
}

func (m *Status) Handler(ctx context.Context, l logger.CategoryLogger, w http.ResponseWriter, r *http.Request) {
	//uuid := ctx.Value(rpc.CtxUuidKey).(string)

}
