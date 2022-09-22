package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MelkoV/go-learn-api/api/user"
	"github.com/MelkoV/go-learn-api/middleware"
	"github.com/MelkoV/go-learn-api/rpc"
	"github.com/MelkoV/go-learn-logger/logger"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
)

func makeUuid() string {
	id := uuid.New()
	return id.String()
}

type Server struct {
	l logger.CategoryLogger
}

func NewApiServer(l logger.CategoryLogger) *Server {
	return &Server{l: l}
}

func Serve(port int, l logger.CategoryLogger) {
	mux := http.NewServeMux()
	server := NewApiServer(l)

	mux.HandleFunc("/user", server.userHandler)

	l.Info("running API server on port %d", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}

func prepareAction(w http.ResponseWriter, r *http.Request, l logger.CategoryLogger) (context.Context, string, error) {
	ctx := context.Background()
	id := makeUuid()
	ctx = context.WithValue(ctx, rpc.CtxUuidKey, id)
	l.WithUuid(id)
	if err := middleware.WithPostData(w, r); err != nil {
		return nil, "", err
	}
	method := r.Header.Get(rpc.MethodHeader)
	if method == "" {
		return nil, "", errors.New("empty method")
	}
	return ctx, method, nil
}

func runAction(ctx context.Context, l logger.CategoryLogger, w http.ResponseWriter, r *http.Request, action rpc.Action) {
	body, _ := ioutil.ReadAll(r.Body)
	l.Info("run action with request %v", string(body[:]))
	if err := json.Unmarshal(body, &action); err != nil {
		l.Error("unmarshal error: %v", err)
		return
	}
	action.Handle(ctx, l, w, r)
}

func (s *Server) userHandler(w http.ResponseWriter, r *http.Request) {
	l := s.l.AddSubCategory("user")
	ctx, method, err := prepareAction(w, r, l)
	if err != nil {
		l.Error("can't prepare action: %s", err)
		rpc.WriteError(w, rpc.CodeServerError, rpc.MessageServerError)
		return
	}
	l = l.AddSubCategory(method)
	var action rpc.Action
	switch method {
	case "login":
		action = &user.LoginRequest{}
	default:
		l.Error("not found handler for method %s", method)
		rpc.WriteError(w, rpc.CodeNotFound, rpc.MessageNotFound)
		return
	}
	runAction(ctx, l, w, r, action)
}
