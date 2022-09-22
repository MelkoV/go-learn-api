package service

import (
	"github.com/MelkoV/go-learn-proto/proto/user"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	conn, err := grpc.Dial(viper.GetString("services.user"), opts...)
	if err != nil {
		return nil, err
	}

	client := user.NewUserServiceClient(conn)
	return client, nil
}
