package server

import (
	"context"

	"github.com/Luiggy102/go-grpc-protobuf/models"
	"github.com/Luiggy102/go-grpc-protobuf/repository"
	"github.com/Luiggy102/go-grpc-protobuf/studentpb"
)

type StudentServer struct {
	repo repository.Repository
	studentpb.UnimplementedStudentServiceServer
}

func NewStudentServer(repo repository.Repository) *StudentServer {
	return &StudentServer{repo: repo}
}

func (s *StudentServer) GetStudent(ctx context.Context, req *studentpb.GetStudentRequest) (*studentpb.Student, error) {
	// find the student by id
	student, err := s.repo.GetStudent(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &studentpb.Student{
		Id:   student.Id,
		Name: student.Name,
		Age:  student.Age,
	}, nil
}

func (s *StudentServer) SetStudent(ctx context.Context, req *studentpb.Student) (*studentpb.SetStudentResponse, error) {

	// add to the db
	student := &models.Student{
		Id:   req.GetId(),
		Name: req.GetName(),
		Age:  req.GetAge(),
	}

	err := s.repo.SetStudent(ctx, *student)
	if err != nil {
		return nil, err
	}

	// return the id
	return &studentpb.SetStudentResponse{
		Id: req.Id,
	}, nil

}
