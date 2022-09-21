package api

import (
	"context"
	"fmt"
	"github.com/MelkoV/go-learn-api/api/user"
	"github.com/MelkoV/go-learn-api/middleware"
	"github.com/MelkoV/go-learn-api/rpc"
	"github.com/MelkoV/go-learn-common/app"
	"github.com/MelkoV/go-learn-logger/logger"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func makeUuid() string {
	id := uuid.New()
	return id.String()
}

type Server struct {
	l *logger.CategoryLogger
}

func NewApiServer(l *logger.CategoryLogger) *Server {
	return &Server{l: l}
}

func Serve(port int, l *logger.CategoryLogger) {
	mux := http.NewServeMux()
	server := NewApiServer(l)

	mux.HandleFunc("/user", server.userHandler)

	l.Format("init", app.SYSTEM_UUID, "running API server on port %d", port).Info()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}

func prepareAction(category string, w http.ResponseWriter, r *http.Request, l *logger.CategoryLogger) (context.Context, rpc.JsonRequest, error) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, rpc.CtxUuidKey, makeUuid())
	var data rpc.JsonRequest
	if err := middleware.WithPostData(w, r, &data); err != nil {
		return nil, rpc.JsonRequest{}, err
	}
	l.Format(category, ctx.Value(rpc.CtxUuidKey).(string), "incoming request: %v", data).Info()
	return ctx, data, nil
}

func runAction(ctx context.Context, data rpc.JsonRequest, a rpc.Action, l *logger.CategoryLogger, w http.ResponseWriter, r *http.Request) error {
	if err := rpc.FillParams(data, a); err != nil {
		return err
	}
	l.Format("user", ctx.Value(rpc.CtxUuidKey).(string), "start action %s with data %v", data.Method, a).Info()
	a.Handler(ctx, l, w, r)
	return nil
}

func (s *Server) userHandler(w http.ResponseWriter, r *http.Request) {
	ctx, data, err := prepareAction("user", w, r, s.l)
	if err != nil {
		s.l.Format("user", app.SYSTEM_UUID, "middleware error: %v", err).Error()
		return
	}
	id := ctx.Value(rpc.CtxUuidKey).(string)

	var a rpc.Action

	if data.Method == "login" {
		a = &user.Login{}
	} else {
		s.l.Format("user", id, "no action for method: %v", data.Method).Error()
		return
	}

	if err = runAction(ctx, data, a, s.l, w, r); err != nil {
		s.l.Format("user", id, "run action error: %v", err).Error()
		return
	}

	/*a := rpc.NewJsonAction(data.Method)

	if a.Method == "login" {
		var model user.LoginRequest
		err = rpc.FillParams(data, &model)
		a.Register(user.LoginHandler)
	}*/

	/*
		type call struct {
			method  string
			handler func(l *logger.CategoryLogger, v interface{})
		}
		var callable = call{
			method: data.Method,
		}

		if callable.method == "login" {
			var model user.LoginRequest
			err = rpc.FillParams(data, &model)
			callable.handler = //func() {
				user.LoginHandler
			//}
		} else {
			s.l.Format("user", id, "no action for method: %v", callable.method).Error()
			return
		}

		if err != nil {
			s.l.Format("user", id, "fill params error: %v", err).Error()
			return
		}

		s.l.Format("user", id, "start action %v", callable.method).Error()

		callable.handler()

	*/
}
