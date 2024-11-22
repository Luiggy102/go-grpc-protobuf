package server

import (
	"context"
	"io"

	"github.com/Luiggy102/go-grpc-protobuf/models"
	"github.com/Luiggy102/go-grpc-protobuf/repository"
	"github.com/Luiggy102/go-grpc-protobuf/testpb"
)

// grpc
type TestServer struct {
	repo repository.Repository
	testpb.UnimplementedTestServiceServer
}

// constructor
func NewTestServer(repo repository.Repository) *TestServer {
	return &TestServer{repo: repo}
}

// service TestService {
//   rpc GetTest(GetTestRequest) returns (Test);
//   rpc SetTest(Test) returns (SetTestResponse);
// }

// implementation grpc
func (s *TestServer) GetTest(ctx context.Context, req *testpb.GetTestRequest) (*testpb.Test, error) {
	test, err := s.repo.GetTest(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &testpb.Test{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}
func (s *TestServer) SetTest(ctx context.Context, req *testpb.Test) (*testpb.SetTestResponse, error) {
	test := &models.Test{
		Id:   req.GetId(),
		Name: req.GetName(),
	}
	err := s.repo.SetTest(ctx, *test)
	if err != nil {
		return nil, err
	}
	return &testpb.SetTestResponse{
		Id:   req.Id,
		Name: req.Name,
	}, nil
}

// rpc client streaming
//
//	service TestService {
//	  rpc SetQuestion(stream Question) returns (SetQuestionResponse);
//	}
func (s *TestServer) SetQuestion(stream testpb.TestService_SetQuestionServer) error {
	for {
		// recibe the msg
		msg, err := stream.Recv()
		if err == io.EOF { // the client stops
			// send the response ok
			return stream.SendAndClose(&testpb.SetQuestionResponse{Ok: true})
		}
		if err != nil {
			return nil
		}
		// send the msg to the db
		err = s.repo.SetQuestion(context.Background(), &models.Question{
			Id:       msg.GetId(),
			Question: msg.GetQuestion(),
			Answer:   msg.GetAnswer(),
			TestId:   msg.GetTestId(),
		})
		// if any error send a ok = false response
		if err != nil {
			return stream.SendAndClose(&testpb.SetQuestionResponse{Ok: false})
		}
	}
}
