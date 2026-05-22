package main

import (
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
