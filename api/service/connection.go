package service

import (
	"context"
	"github.com/MelkoV/go-learn-proto/proto/user"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

var userService user.UserServiceClient

func UserServiceClient() (user.UserServiceClient, error) {
	if userService == nil {
		var err error
		userService, err = newUserServiceClient()
		if err != nil {
			return nil, err
		}
	}

	return userService, nil
}

func newUserServiceClient() (user.UserServiceClient, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//opts = append(opts, grpc.WithTimeout())

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	conn, err := grpc.DialContext(ctx, viper.GetString("services.user"), opts...)
	if err != nil {
		return nil, err
	}

	client := user.NewUserServiceClient(conn)
	return client, nil
}
