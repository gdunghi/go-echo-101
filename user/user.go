package user

import (
	"database/sql"
	"log"
)

type (
	UserModelInterface interface {
		GetUserByID(id int) User
		GetAll() ([]User, error)
		Create(u User) (int64, error)
	}

	User struct {
		ID       int    `json:"id" db:"id"`
		Username string `json:"username" db:"username"`
		Password string `json:"password" db:"password"`
	}

	UserModel struct {
		db *sql.DB
	}
)

//NewUserModel ... new NewUserModel
func NewUserModel(db *sql.DB) *UserModel {
	return &UserModel{
		db: db,
	}
}

//GetUserByID ...
func (u *UserModel) GetUserByID(id int) User {
	user := User{}
	rows, err := u.db.Query("select id, username from users where id=?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(user.ID, user.Username)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return user
}

//GetAll ... get all user from DB
func (u *UserModel) GetAll() ([]User, error) {
	users := []User{}
	rows, err := u.db.Query(`select * from users`)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.ID, &u.Username, &u.Password)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		log.Println(u.ID, u.Username)
		users = append(users, u)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return users, nil
}

//Create ... Create user with user struc
// return int64
func (u *UserModel) Create(user User) (int64, error) {

	res, err := u.db.Exec("INSERT INTO users(username, password) VALUES (?, ?)", user.Username, user.Password)

	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
