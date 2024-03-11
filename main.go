package main

import (
	"ToDoApplication/handlers"
	"github.com/CloudyKit/jet/v6"
	"github.com/labstack/echo/v4"
	"io"
	"log"
)

func main() {

	t := &Template{View: jet.NewSet(jet.NewOSFileSystemLoader("./templates"), jet.InDevelopmentMode())}
	e := echo.New()
	e.Renderer = t
	e.Static("/templates", "templates") //static folder

	e.GET("/", handlers.LoginPage)
	e.POST("/", handlers.LoginHandler)
	e.GET("/signup", handlers.SignupPage)
	e.POST("/signup", handlers.CreateUserHandler)
	e.GET("/home", handlers.HomeHandler)
	e.GET("/logout", handlers.LogoutHandler)

	e.HTTPErrorHandler = handlers.HTTPErrorHandler

	log.Fatal(e.Start(":1323"))
}

type Template struct {
	View *jet.Set
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	template, err := t.View.GetTemplate(name)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	vars := make(jet.VarMap)
	if c.Get("flash") != nil {
		vars.Set("flash", c.Get("flash"))
	}
	err = template.Execute(w, vars, data)
	if err != nil {
		return err
	}
	return nil
}
