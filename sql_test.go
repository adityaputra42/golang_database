package golangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
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

func TestSqlInjection(t *testing.T) {

	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "admin' #"
	password := "salah"
	script := "SELECT username FROM user where username ='" + username + "'and password ='" + password + "' limit 1"
	rows, err := db.QueryContext(ctx, script)
	fmt.Println(script)
	if err != nil {
		panic(err)
	}
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses login =>", username)
	} else {
		fmt.Println("Gagal login ")
	}

	defer rows.Close()
}

func TestSqlInjectionSafe(t *testing.T) {

	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "admin"
	password := "admin"
	script := "SELECT username FROM user where username = ? and password = ? limit 1"
	rows, err := db.QueryContext(ctx, script, username, password)
	fmt.Println(script)
	if err != nil {
		panic(err)
	}
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses login =>", username)
	} else {
		fmt.Println("Gagal login ")
	}

	defer rows.Close()
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "gita72"
	pasword := "123456"
	script := "INSERT INTO user(username, password) VALUES(?,?)"
	_, err := db.ExecContext(ctx, script, username, pasword)
	if err != nil {
		panic(err)
	}
	fmt.Println("Succes insert new user")
}

func TestExecSqlAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	email := "gita21@gmail.com"
	coment := "bolo-bolo"
	script := "INSERT INTO comments(email, coment) VALUES(?,?)"
	result, err := db.ExecContext(ctx, script, email, coment)
	if err != nil {
		panic(err)
	}
	insertId, error := result.LastInsertId()
	if error != nil {
		panic(error)
	}
	fmt.Println("Last Insert id :", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments(email, coment) VALUES(?,?)"
	stmt, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "aditya" + strconv.Itoa(i) + "@gmail.com"
		coment := "Komentar saya ke" + strconv.Itoa(i)
		result, err := stmt.ExecContext(ctx, email, coment)
		if err != nil {
			panic(err)
		}
		insertId, error := result.LastInsertId()
		if error != nil {
			panic(error)
		}
		fmt.Println("Comment Id :", insertId)
	}

}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	script := "INSERT INTO comments(email, coment) VALUES(?,?)"
	for i := 0; i < 10; i++ {
		email := "gitaPrigi" + strconv.Itoa(i) + "@gmail.com"
		coment := "Gita Prigi ke =>" + strconv.Itoa(i)
		result, err := tx.ExecContext(ctx, script, email, coment)
		if err != nil {
			panic(err)
		}
		insertId, error := result.LastInsertId()
		if error != nil {
			panic(error)
		}
		fmt.Println("Comment Id :", insertId)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}

}
