package main

import (
	"context"
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
		log.Fatalf("could not connect %v\n", err)
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
	// get test
	req := &testpb.GetTestRequest{Id: "t1"}

	res, err := c.GetTest(context.Background(), req)
	if err != nil {
		log.Fatalln("error while calling GetTest;", err)
	}

	log.Println("response from server:", res)
}
func DoClientStreaming(c testpb.TestServiceClient) {
	// set questions

}
func DoServerStreaming(c testpb.TestServiceClient) {
	// get student per test

}
func DoBidirecitonalStreaming(c testpb.TestServiceClient) {
	// take test

}
