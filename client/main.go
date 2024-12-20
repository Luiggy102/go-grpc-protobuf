package main

import (
	"context"
	"io"
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

	// DoUnary(c)
	// DoClientStreaming(c)
	// DoServerStreaming(c)
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
		{Id: "q14", Question: "question 14", Answer: "answer 14", TestId: "t1"},
		{Id: "q15", Question: "question 15", Answer: "answer 15", TestId: "t1"},
		{Id: "q16", Question: "question 16", Answer: "answer 16", TestId: "t1"},
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
	req := &testpb.GetStudentPerTestRequest{TestId: "t1"}

	stream, err := c.GetStudentPerTest(context.Background(), req)
	if err != nil {
		log.Fatalln("error while calling GetStudentPerTest", err)
	}

	// catch the server response
	for {
		msg, err := stream.Recv()
		if err == io.EOF { // completed
			break
		}
		if err != nil {
			log.Fatalln("error while receiving response:", err)
		}
		log.Println("response from server:", msg)
	}
}
func DoBidirecitonalStreaming(c testpb.TestServiceClient) {
	// take test

	req := &testpb.TakeTestRequest{Answer: "answer"}

	numberOfQuestion := 4
	waitchannel := make(chan struct{})

	stream, err := c.TakeTest(context.Background())
	if err != nil {
		log.Fatalln("error while calling TakeTest", err)
	}

	// client stremaing goroutine (send the answers)
	go func() {
		for i := 0; i < numberOfQuestion; i++ {
			err = stream.Send(req)
			if err != nil {
				log.Fatalln("error while sending stremaing:", err)
				break
			}
			time.Sleep(time.Second * 1)
		}
	}()

	// server stremaing goroutine (undefined responses)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln("error while receiving response:", err)
				break
			}
			log.Println("response from server:", res)
		}
		close(waitchannel)
	}()
	<-waitchannel

}
