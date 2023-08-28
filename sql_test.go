package golangdatabase

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
 script := "INSERT INTO customer(id, name) VALUES('gita','Gita Prigi Andani')"
	_, err := db.ExecContext(ctx,script)
if err != nil {
	panic(err)
}
fmt.Println("Succes insert new customer")
}
