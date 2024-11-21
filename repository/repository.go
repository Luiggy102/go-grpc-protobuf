package repository

import (
	"context"

	"github.com/Luiggy102/go-grpc-protobuf/models"
)

// Repository pattern
type Repository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student models.Student) error
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
