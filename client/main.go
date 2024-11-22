package main

import (
	"log"

	"github.com/Luiggy102/go-grpc-protobuf/testpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// create grpc conn
	// cc, err := grpc.Dial("localhost:5070",
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()))

	// grpc conn to the server
	cc, err := grpc.NewClient("localhost:5070", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect %v\n", err)
	}
	// close the conn at the end
	defer cc.Close()

	// new client
	c := testpb.NewTestServiceClient(cc)

	// different grpc methods
	DoUnary(c)
	DoClientStreaming(c)
	DoServerStreaming(c)
	DoBidirecitonalStreaming(c)
}

func DoUnary(c testpb.TestServiceClient) {

}
func DoClientStreaming(c testpb.TestServiceClient) {

}
func DoServerStreaming(c testpb.TestServiceClient) {

}
func DoBidirecitonalStreaming(c testpb.TestServiceClient) {

}
