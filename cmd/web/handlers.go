package main

import (
	"io"
	"log"
	"net/http"
)

func (app *application) getTestPage(w http.ResponseWriter, r *http.Request) {

	data := app.PageContent

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
		log.Print(err.Error())
		return //Return because if resp is nil resp.Body.Close() will panic
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		w.Write([]byte("Status code not OK"))
		log.Print("resp status code not OK")
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return
	}
}
