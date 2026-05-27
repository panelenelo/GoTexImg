package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /images/{img}", app.getImgProxy)
	mux.HandleFunc("GET /testpage/{id}", app.getTestPage)
	mux.HandleFunc("GET /teststatic", app.getTestStatic)

	return mux
}
