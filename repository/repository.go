package repository

import (
	"context"

	"github.com/Luiggy102/go-grpc-protobuf/models"
)

// Repository pattern
type Repository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student models.Student) error
	GetTest(ctx context.Context, id string) (*models.Test, error)
	SetTest(ctx context.Context, test models.Test) error
	SetQuestion(ctx context.Context, question *models.Question) error
	EnrollStudents(ctx context.Context, enrollment *models.Enrollment) error
	GetStudentPerTest(ctx context.Context, testId string) ([]*models.Student, error)
	GetQuestionsPerTest(ctx context.Context, testId string) ([]*models.Question, error)
}

var implementation Repository

// dependency infection
func SetRepository(repository Repository) {
	implementation = repository
}

func GetStudent(ctx context.Context, id string) (*models.Student, error) {
	return implementation.GetStudent(ctx, id)
}
func SetStudent(ctx context.Context, student models.Student) error {
	return implementation.SetStudent(ctx, student)
}

func GetTest(ctx context.Context, id string) (*models.Test, error) {
	return implementation.GetTest(ctx, id)
}
func SetTest(ctx context.Context, test models.Test) error {
	return implementation.SetTest(ctx, test)
}

func SetQuestion(ctx context.Context, question *models.Question) error {
	return implementation.SetQuestion(ctx, question)
}

func EnrollStudents(ctx context.Context, enrollment *models.Enrollment) error {
	return implementation.EnrollStudents(ctx, enrollment)
}
func GetStudentPerTest(ctx context.Context, testId string) ([]*models.Student, error) {
	return implementation.GetStudentPerTest(ctx, testId)
}

func GetQuestionsPerTest(ctx context.Context, testId string) ([]*models.Question, error) {
	return implementation.GetQuestionsPerTest(ctx, testId)
}
