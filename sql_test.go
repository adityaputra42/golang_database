package golangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO customer(id, name) VALUES('1','Paijo')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Succes insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id : ", id)
		fmt.Println("Name : ", name)
	}
	defer rows.Close()

}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id, name , email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id : ", id)
		fmt.Println("Name : ", name)
		if email.Valid {
			fmt.Println("email : ", email.String)
		} else {
			fmt.Println("email : Null")
		}
		fmt.Println("Balance : ", balance)
		fmt.Println("Rating : ", rating)
		if birthDate.Valid {
			fmt.Println("Birthdate : ", birthDate.Time)
		} else {
			fmt.Println("Birthdate : Null")
		}
		fmt.Println("Married : ", married)
		fmt.Println("CreatedAt : ", createdAt)
		fmt.Println("<===============Next==============>")
	}
}
