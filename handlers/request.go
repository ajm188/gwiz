package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

type Request interface {
	Render(string)
	RenderFile(string)
	RenderTemplate(string, interface{})

	Redirect(string, int)

	Error(int, error)
}

type GwizRequest struct {
	http.ResponseWriter
	*http.Request
}

func (g *GwizRequest) Render(raw string) {
	fmt.Fprintf(g.ResponseWriter, raw)
}

func (g *GwizRequest) RenderFile(filename string) {
	http.ServeFile(g.ResponseWriter, g.Request, filename)
}

func (g *GwizRequest) Redirect(url string, status int) {
	http.Redirect(g.ResponseWriter, g.Request, url, status)
}

func (g *GwizRequest) RenderTemplate(templatePath string, data interface{}) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		g.Error(500, err)
		return
	}

	err = t.Execute(g.ResponseWriter, data)
	if err != nil {
		g.Error(500, err)
		return
	}
}

func (g *GwizRequest) Error(status int, err error) {
	// TODO: set the status code
	fmt.Fprintf(g.ResponseWriter, "Got error: %s\n", err)
}
