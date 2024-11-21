package server

import (
	"context"

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
	student, err := s.repo.GetTest(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &testpb.Test{
		Id:   student.Id,
		Name: student.Name,
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
		Id: req.Id,
	}, nil
}
