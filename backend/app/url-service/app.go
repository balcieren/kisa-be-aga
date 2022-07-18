package urlservice

import (
	"backend/app/url-service/database"
	"backend/pkg/config"
	"backend/pkg/proto"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func Start() {
	c := config.New()

	db, err := database.NewPostgreSQL(c.UrlServiceDatabaseUrl)
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", ":"+c.UrlServicePort)

	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	proto.RegisterUrlServiceServer(srv, &Service{
		db: db,
	})

	println(fmt.Sprintf("Service has started on :%s", c.UrlServicePort))
	panic(srv.Serve(lis))
}
