package http

import (
	"net/http"

	. "github.com/x07-it/gomps"
	ghttp "github.com/x07-it/gomps/http"

	"app/html"
)

func Home(mux *http.ServeMux) {
	mux.Handle("GET /", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		// Let's pretend this comes from a db or something.
		items := []string{"Super", "Duper", "Nice"}
		return html.HomePage(items), nil
	}))
}

func About(mux *http.ServeMux) {
	mux.Handle("GET /about", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		return html.AboutPage(), nil
	}))
}
