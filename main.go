package main

import (
	"errors"
	"fmt"

	"github.com/skhut/go-sql-client/sqlclient"
)

const (
	sqlGetUser = "SELECT customer_id, email FROM customer WHERE customer_id=?;"
)

var (
	dbClient sqlclient.SqlClient
)

type User struct {
	Id    int64
	Email string
}

func init() {
	var err error
	sqlclient.StartMockServer()
	dbClient, err = sqlclient.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root", "mysql@local4admin", "127.0.0.1:3306", "sakila"))
	if err != nil {
		panic(err)
	}
}

func main() {
	user, err := GetUser(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(user.Id)
	fmt.Println(user.Email)
}

func GetUser(id int64) (*User, error) {
	sqlclient.AddMock(sqlclient.Mock{
		Query: sqlGetUser,
		Args:  []interface{}{1},
		//Error: errors.New("error creating query"),
		Columns: []string{"id", "email"},
		Rows: [][]interface{}{
			{1, "email1"},
			{2, "email2"},
		},
	})
	var user User
	rows, err := dbClient.Query(sqlGetUser, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.HasNext() {
		if err = rows.Scan(&user.Id, &user.Email); err != nil {
			return nil, err
		}
		fmt.Printf("User:%#v\n", &user)
		return &user, nil
	}
	return nil, errors.New("user not found")
}
