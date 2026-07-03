package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.getHome)
	mux.HandleFunc("GET /images/{img}", app.getImgProxy)
	mux.HandleFunc("GET /texts/{id}", app.getTextView)
	mux.HandleFunc("GET /testpage/{id}", app.getTestPage)
	mux.HandleFunc("GET /teststatic", app.getTestStatic)
	mux.HandleFunc("GET /create", app.getCreateText)
	mux.HandleFunc("POST /create", app.postCreateText)

	return mux
}
