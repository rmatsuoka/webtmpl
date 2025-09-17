package xsql_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/rmatsuoka/webtmpl/internal/x/xsql"
)

func ExampleQuery() {
	type User struct {
		ID   int64
		Name string
	}

	var db *sql.DB

	var users []*User

	rows := xsql.Query(context.Background(), db, `select id, name from users`)
	defer rows.Close()

	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Name)
		users = append(users, &u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(users)
}
