package main

import (
	"log"
	"net"

	"github.com/Luiggy102/go-grpc-protobuf/database"
	"github.com/Luiggy102/go-grpc-protobuf/server"
	"github.com/Luiggy102/go-grpc-protobuf/studentpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// port
	listener, err := net.Listen("tcp", ":5060")
	if err != nil {
		log.Fatal(err)
	}

	// new db
	repo, err := database.NewPostgresRepo(database.PgUrl)
	if err != nil {
		log.Fatal(err)
	}

	// new server
	server := server.NewStudentServer(repo)

	// gRpc connection
	s := grpc.NewServer()

	// register StudentServer
	studentpb.RegisterStudentServiceServer(s, server)

	// add reflection
	reflection.Register(s)

	// run
	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
