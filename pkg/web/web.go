package web

import (
	"html/template"
	"io"
	"net/http"
	"vitt/pkg/store"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/urfave/cli/v2"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Init(ctx *cli.Context) error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	t := &Template{
		templates: template.Must(template.ParseGlob("pkg/web/views/*.html")),
	}
	e.Renderer = t

	// Routes
	e.GET("/", hello)
	e.GET("/transactions", func(c echo.Context) error {
		return getTransactions(c, ctx)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

	return nil
}

// Handler
func hello(c echo.Context) error {
	return c.Render(http.StatusOK, "main", "")
}

func getTransactions(c echo.Context, ctx *cli.Context) error {
	db := ctx.Context.Value("db").(*store.Store)

	transactions, err := db.GetTransactions()
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "transactions", transactions)
}
