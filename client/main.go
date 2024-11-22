package main

import (
	"context"
	"log"
	"time"

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
	questions := []*testpb.Question{
		{Id: "q11", Question: "question 11", Answer: "answer 11", TestId: "t1"},
		{Id: "q12", Question: "question 12", Answer: "answer 12", TestId: "t1"},
		{Id: "q13", Question: "question 13", Answer: "answer 13", TestId: "t1"},
	}

	stream, err := c.SetQuestion(context.Background())
	if err != nil {
		log.Fatalln("error while calling SetQuestion", err)
	}

	// send the stream
	for _, q := range questions {
		log.Println("sending question:", q.Id)
		err = stream.Send(q)
		if err != nil {
			log.Fatalln("error while stream", err)
		}
		time.Sleep(time.Second * 2)
	}

	// close the coneciton and get the answer
	msg, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("error while receiving response:", err)
	}

	log.Println("response from server:", msg)

}
func DoServerStreaming(c testpb.TestServiceClient) {
	// get student per test

}
func DoBidirecitonalStreaming(c testpb.TestServiceClient) {
	// take test

}
