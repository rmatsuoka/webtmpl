package xsql

import (
	"context"
	"database/sql"
	"fmt"
)

func ExampleQuery() {
	type User struct {
		ID   int64
		Name string
	}

	var db *sql.DB

	var (
		users []*User
		err   error
	)
	for scan, scanErr := range Query(context.Background(), db, `select id, name from users`) {
		if scanErr != nil {
			err = scanErr
			break
		}
		var u User
		scan(&u.ID, &u.Name)
		users = append(users, &u)
	}

	fmt.Println(users, err)
}
