package main

import (
	"html/template"
	"io"

	"./handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Renderer for templates
type Renderer struct {
	templates *template.Template
}

// Render templates
func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return r.templates.ExecuteTemplate(w, name, data)
}

func main() {
	echoServer := echo.New()

	// Set for debug.
	echoServer.Use(middleware.Logger())
	echoServer.Use(middleware.Recover())

	// Set rendering config.
	// Specially, change the default deliminators to avoid confliction Template with Vue.js.
	templ := template.New("").Delims("[[", "]]")
	echoServer.Renderer = &Renderer{
		templates: template.Must(templ.ParseGlob("views/*.html")),
	}

	// Set routing.
	echoServer.GET("/", handler.CalculatorGET())
	echoServer.POST("/", handler.CalculatorPOST())

	// Set resource path.
	echoServer.Static("/public", "public")

	// Start this echo server.
	echoServer.Start(":8080")
}
