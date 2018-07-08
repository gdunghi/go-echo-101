package main

import (
	"github.com/gdunghi/go-echo-101/db"
	"github.com/gdunghi/go-echo-101/user"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

	d := db.DBConnect()

	uh := user.NewHandler(user.NewUserModel(d))

	e.GET("/users/:id", uh.GetUserByID)
	e.GET("/users", uh.GetAllUsers)

	e.Logger.Fatal(e.Start(":1324"))
}
