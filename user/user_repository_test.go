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

	um := NewUserRepository(db)
	u := um.GetUserByID(1)
	expect := User{
		ID:       1,
		Username: "foobar",
	}
	assert.Equal(t, expect, u)

}

func TestCreateUser(t *testing.T) {
	//with golang test without assert test framework
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	result := sqlmock.NewResult(1, 1)
	mock.ExpectExec("^INSERT INTO users").
		WithArgs("foo", "foopass").
		WillReturnResult(result)

	//res, err := db.Exec("INSERT INTO users (username,password) VALUES (?,?)", "foo", "foopass")
	um := NewUserRepository(db)
	id, err := um.Create(User{Username: "foo", Password: "foopass"})

	if id != 1 {
		t.Errorf("expected last insert id to be 1, but got %d instead", id)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}
