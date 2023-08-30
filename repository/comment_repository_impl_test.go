package repository

import (
	"context"
	"fmt"
	golangdatabase "golang_database"
	"golang_database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInsertComment(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())
	ctx := context.Background()

	comment := entity.Comment{
		Email:   "repository24@test.com",
		Comment: "Belajar golang database",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

}

func TestFindByIdComment(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())
	ctx := context.Background()
	comment, err := commentRepository.FindById(ctx, 12)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAllComment(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())
	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)

	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}

}
