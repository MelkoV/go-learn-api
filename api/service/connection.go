package service

import (
	"context"
	"github.com/MelkoV/go-learn-api/rpc"
	"github.com/MelkoV/go-learn-logger/logger"
	"github.com/MelkoV/go-learn-proto/proto/user"
	pb "github.com/MelkoV/go-learn-proto/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserServiceClient(ctx context.Context, l *logger.CategoryLogger) user.UserServiceClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("localhost:5001", opts...)
	if err != nil {
		l.Format("service/user", ctx.Value(rpc.CtxUuidKey).(string), "fail to dial: %v", err).Fatal()
	}

	client := pb.NewUserServiceClient(conn)
	return client
}
