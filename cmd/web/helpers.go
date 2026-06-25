package main

import (
	"bytes"
	"html/template"
	"net/http"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, files []string, data PageContent) {
	ts, err := template.New("Tmplt").Funcs(tmpltFunctions).ParseFiles(files...)
	if err != nil {
		app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	buff := new(bytes.Buffer)
	err = ts.ExecuteTemplate(buff, "base", data)
	if err != nil {
		app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(status)

	buff.WriteTo(w)
}

func (app *application) renderTest(w http.ResponseWriter, r *http.Request, status int, files []string, data PageContent) {
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	buff := new(bytes.Buffer)
	err = ts.ExecuteTemplate(buff, "testbase", data)
	if err != nil {
		app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(status)

	buff.WriteTo(w)

}
