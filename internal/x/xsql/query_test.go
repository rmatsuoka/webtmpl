package xsql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

func ExampleQuery() {
	type User struct {
		ID   int64
		Name string
	}

	var db *sql.DB

	var users []*User

	rows := Query(context.Background(), db, `select id, name from users`)
	for scan := range rows.ScanSeq() {
		var u User
		scan(&u.ID, &u.Name)
		users = append(users, &u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(users)
}
