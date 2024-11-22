package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/Luiggy102/go-grpc-protobuf/models"
	_ "github.com/lib/pq"
)

var PgUrl = "postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable"

type PostgresRepo struct {
	Db *sql.DB
}

func NewPostgresRepo(url string) (*PostgresRepo, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepo{Db: db}, nil
}

func (repo *PostgresRepo) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	rows, err := repo.Db.QueryContext(ctx, "select id, name, age from students where id = $1", id)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	s := models.Student{}
	for rows.Next() {
		err := rows.Scan(&s.Id, &s.Name, &s.Age)
		if err != nil {
			return nil, err
		}
	}
	return &s, nil
}
func (repo *PostgresRepo) SetStudent(ctx context.Context, student models.Student) error {
	_, err := repo.Db.ExecContext(ctx,
		"insert into students (id, name,age) values ($1,$2,$3)",
		student.Id, student.Name, student.Age)
	return err
}

func (repo *PostgresRepo) GetTest(ctx context.Context, id string) (*models.Test, error) {
	rows, err := repo.Db.QueryContext(ctx, "select id, name from tests where id = $1", id)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	t := models.Test{}
	for rows.Next() {
		err := rows.Scan(&t.Id, &t.Name)
		if err != nil {
			return nil, err
		}
	}
	return &t, nil

}
func (repo *PostgresRepo) SetTest(ctx context.Context, test models.Test) error {
	_, err := repo.Db.ExecContext(ctx,
		"insert into tests (id, name) values ($1,$2)",
		test.Id, test.Name)
	return err
}

func (repo *PostgresRepo) SetQuestion(ctx context.Context, question *models.Question) error {
	_, err := repo.Db.ExecContext(ctx,
		"insert into questions (id, question, answer, test_id) values ($1,$2,$3,$4)",
		question.Id, question.Question, question.Answer, question.TestId)
	return err
}
