package main

import (
	"ToDoApplication/handlers"
	"github.com/CloudyKit/jet/v6"
	"github.com/labstack/echo/v4"
	"io"
	"log"
)

func main() {
	// Template initializeren.
	t := &Template{View: jet.NewSet(jet.NewOSFileSystemLoader("./templates"), jet.InDevelopmentMode())}

	// Echo initializeren.
	e := echo.New()
	e.Renderer = t
	e.Static("/templates", "templates")

	// Routes.
	e.GET("/", handlers.HomePage)
	e.GET("/login", handlers.LoginPage)
	e.POST("/login", handlers.LoginHandler)
	e.GET("/signup", handlers.SignupPage)
	e.POST("/signup", handlers.CreateUserHandler)
	e.GET("/logout", handlers.LogoutHandler)
	e.GET("/todos", handlers.ToDosHandler)
	e.GET("/todos/new", handlers.NewToDoPageHandler)
	e.POST("/todos/new", handlers.CreateToDoHandler)
	e.GET("/todos/delete/:id", handlers.DeleteToDoHandler)

	e.HTTPErrorHandler = handlers.HTTPErrorHandler

	if err := e.Start(":1323"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

type Template struct {
	View *jet.Set
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	log.Printf("Rendering template: %s\n", name)

	// Template halen.
	template, err := t.View.GetTemplate(name)
	if err != nil {
		log.Printf("Failed to retrieve template %s: %s\n", name, err.Error())
		return err
	}

	// Variables preparen.
	vars := make(jet.VarMap)
	if c.Get("flash") != nil {
		vars.Set("flash", c.Get("flash"))
	}

	// Templates executen.
	log.Printf("Executing template: %s\n", name)
	err = template.Execute(w, vars, data)
	if err != nil {
		log.Printf("Failed to execute template %s: %s\n", name, err.Error())
		return err
	}

	log.Printf("Template %s rendered successfully\n", name)
	return nil
}
