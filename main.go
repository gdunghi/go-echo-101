package main

import (
	"github.com/gdunghi/go-echo-101/db"
	"github.com/gdunghi/go-echo-101/user"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

	c := db.DBConnect()

	uh := user.NewHandler(user.NewUserRepository(c))

	e.GET("/users/:id", uh.GetUserByID)
	e.GET("/users", uh.GetAll)

	e.POST("/users", uh.Create)
	e.Logger.Fatal(e.Start(":1324"))
}
