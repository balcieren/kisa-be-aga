package service

import (
	"backend/pkg/config"
	"backend/pkg/proto"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct {
	Url proto.UrlServiceClient
}

func New(c *config.Config) (*Service, error) {
	conn, err := grpc.Dial(fmt.Sprintf("url-svc:%s", c.UrlServicePort), grpc.WithTransportCredentials(insecure.NewCredentials()))

	urlServiceClient := proto.NewUrlServiceClient(conn)

	return &Service{
		Url: urlServiceClient,
	}, err
}
