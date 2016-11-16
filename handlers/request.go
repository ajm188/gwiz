package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

import (
	"github.com/ajm188/gwiz/db"
)

type request interface {
	Render(string)
	RenderFile(string)
	RenderTemplate(string, interface{})

	Redirect(string, int)

	Error(int, error)
}

type Request struct {
	http.ResponseWriter
	*http.Request
	db.Transaction
}

func (r *Request) Render(raw string) {
	fmt.Fprintf(r.ResponseWriter, raw)
}

func (r *Request) RenderFile(filename string) {
	http.ServeFile(r.ResponseWriter, r.Request, filename)
}

func (r *Request) Redirect(url string, status int) {
	http.Redirect(r.ResponseWriter, r.Request, url, status)
}

func (r *Request) RenderTemplate(templatePath string, data interface{}) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		r.Error(500, err)
		return
	}

	err = t.Execute(r.ResponseWriter, data)
	if err != nil {
		r.Error(500, err)
		return
	}
}

func (r *Request) Error(status int, err error) {
	// TODO: set the status code
	fmt.Fprintf(r.ResponseWriter, "Got error: %s\n", err)
}
