package main

import (
	"log"
	"net"

	"github.com/Luiggy102/go-grpc-protobuf/database"
	"github.com/Luiggy102/go-grpc-protobuf/server"
	"github.com/Luiggy102/go-grpc-protobuf/testpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":5070")
	if err != nil {
		log.Fatal(err)
	}
	repo, err := database.NewPostgresRepo(database.PgUrl)
	if err != nil {
		log.Fatal(err)
	}

	// new server
	server := server.NewTestServer(repo)

	// gRpc connection
	s := grpc.NewServer()

	// register StudentServer
	testpb.RegisterTestServiceServer(s, server)

	// add reflection
	reflection.Register(s)

	// run
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
