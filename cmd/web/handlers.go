package main

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func (app *application) getHome(w http.ResponseWriter, r *http.Request) {

	texts, err := app.TModel.Latest()
	if err != nil {
		app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := PageContent{
		Texts: texts,
	}

	files := []string{
		"./ui/html/base.html",
		"./ui/html/pages/home.html",
		"./ui/html/partials/footer.html",
		"./ui/html/partials/nav.html",
	}

	app.render(w, r, http.StatusOK, files, data)

}

func (app *application) getTextView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	text, err := app.TModel.GetTex(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.NotFound(w, r)
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			// log.Printf("Internal Error: %v", err)
			app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		}
		return
	}

	data := PageContent{
		Text: text,
	}

	if text.ImageRefID != 0 {
		img, err := app.TModel.GetImg(int(text.ImageRefID))
		if err != nil {
			// log.Printf("%s", err)
			app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI(), "function", "GetImg")
		} else {
			parts := strings.Split(img.Content, "/")
			filename := parts[len(parts)-1]
			img.Content = "/images/" + filename
			data.ImagesRef = img
		}
	}

	files := []string{
		"./ui/html/base.html",
		"./ui/html/pages/view.html",
		"./ui/html/partials/footer.html",
		"./ui/html/partials/nav.html",
	}

	app.render(w, r, http.StatusOK, files, data)

}

func (app *application) getTestPage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	text, err := app.TModel.GetTex(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.NotFound(w, r)
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			// log.Printf("Internal Error: %v", err)
			app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		}
		return
	}

	data := PageContent{
		Text: text,
	}

	if text.ImageRefID != 0 {
		img, err := app.TModel.GetImg(int(text.ImageRefID))
		if err != nil {
			// log.Printf("%s", err)
			app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI(), "function", "GetImg")
		} else {
			parts := strings.Split(img.Content, "/")
			filename := parts[len(parts)-1]
			img.Content = "/images/" + filename
			data.ImagesRef = img
		}
	}

	files := []string{
		"./ui/html/testbase.html",
		"./ui/html/pages/testpage.html",
		"./ui/html/partials/footer.html",
		"./ui/html/partials/nav.html",
	}

	app.renderTest(w, r, http.StatusOK, files, data)
}

func (app *application) getTestStatic(w http.ResponseWriter, r *http.Request) {
	img := "http://imageserver:80/nepal.jpg"

	resp, err := app.Client.Get(img)
	if err != nil {
		w.Write([]byte(err.Error()))
		app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		return //Return because if resp is nil resp.Body.Close() will panic
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		w.Write([]byte("Status code not OK"))
		// log.Print("resp status code not OK")
		app.Logger.Error("resp status code not OK", "method", r.Method, "uri", r.URL.RequestURI())
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	// log.Printf("Type of resp: %T", resp.Body) is of type *http.cancelTimerBody

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI(), "function", "io.Copy")
		return
	}
}

func (app *application) getImgProxy(w http.ResponseWriter, r *http.Request) {
	url := r.PathValue("img")
	if url == "" {
		// log.Printf("Path empty, no image referenced")
		app.Logger.Error("Path empty, no image referenced", "method", r.Method, "uri", r.URL.RequestURI())
		return
	}

	internalURL := fmt.Sprintf("http://imageserver:80/%s", url)
	resp, err := app.Client.Get(internalURL)
	if err != nil {
		// log.Printf("Error getting image: %s", err.Error())
		app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		resp := fmt.Sprintf("Client.Get(%s)", url)
		app.Logger.Error("Status code not OK", "method", r.Method, "uri", r.URL.RequestURI(), "function", resp)
		return
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "" {
		w.Header().Set("Content-Type", contentType)
	}

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		// log.Printf("Error copying image to response: %v", err)
		app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
	}

}
