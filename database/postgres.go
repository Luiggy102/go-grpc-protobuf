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

func (repo *PostgresRepo) EnrollStudents(ctx context.Context, enrollment *models.Enrollment) error {
	_, err := repo.Db.ExecContext(ctx,
		"insert into enrollments (student_id, test_id) values ($1,$2)",
		enrollment.StudentId, enrollment.TestId)
	return err
}
func (repo *PostgresRepo) GetStudentPerTest(ctx context.Context, testId string) ([]*models.Student, error) {
	rows, err := repo.Db.QueryContext(ctx,
		"SELECT id, name, age FROM students WHERE id IN (SELECT student_id FROM enrollments WHERE test_id = $1)",
		testId)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	students := []*models.Student{}
	for rows.Next() {
		s := models.Student{}
		err := rows.Scan(&s.Id, &s.Name, &s.Age)
		if err != nil {
			return nil, err
		}
		students = append(students, &s)
	}
	return students, nil
}

func (repo *PostgresRepo) GetQuestionsPerTest(ctx context.Context, testId string) ([]*models.Question, error) {
	rows, err := repo.Db.QueryContext(ctx,
		"SELECT id, question FROM question WHERE test_id = $1)",
		testId)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	questions := []*models.Question{}
	for rows.Next() {
		q := models.Question{}
		err = rows.Scan(&q.Id, &q.Question)
		if err != nil {
			return nil, err
		}
		questions = append(questions, &q)
	}
	return questions, nil
}
