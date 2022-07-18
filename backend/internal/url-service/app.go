package urlservice

import (
	"backend/internal/url-service/database"
	"backend/pkg/config"
	"backend/pkg/proto"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func Start() {
	c, err := config.New()
	if err != nil {
		panic(err)
	}

	db, err := database.NewPostgreSQL(c.UrlServiceDatabaseUrl)
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", c.UrlServicePort))

	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	proto.RegisterUrlServiceServer(srv, &Service{
		db: db,
	})

	println(fmt.Sprintf("Service has started on localhost:%s", c.UrlServicePort))
	panic(srv.Serve(lis))
}
