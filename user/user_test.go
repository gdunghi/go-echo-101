package user

import (
	"fmt"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByIDFromDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println("err")
		return
	}

	rows := sqlmock.NewRows([]string{"id", "username"}).AddRow(1, "foobar")
	sql := `select id, username from users where id=?`

	mock.ExpectQuery(sql).WillReturnRows(rows)

	um := SetConnection(db)
	u := um.GetUserByID(1)
	expect := User{
		ID:       1,
		Username: "foobar",
	}
	assert.Equal(t, expect, u)

}

func TestGetUserAllFromDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println("err")
		return
	}

	rows := sqlmock.NewRows([]string{"id", "username", "password"}).AddRow(1, "foo", "pw").AddRow(2, "bar", "pw")
	sql := `select (.*) from users`

	mock.ExpectQuery(sql).WillReturnRows(rows)

	um := SetConnection(db)
	u, _ := um.GetAll()
	expect := []User{{ID: 1, Username: "foo", Password: "pw"}, {ID: 2, Username: "bar", Password: "pw"}}

	assert.Equal(t, expect, u)

}
