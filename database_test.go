package golangdatabase

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDatabase(t *testing.T) {
	fmt.Println("test Connect")
}

func TestConnectDatabase(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Gunakan Database

}
