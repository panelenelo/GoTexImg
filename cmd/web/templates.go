package main

import (
	"GoTexImg/internal/models"
	"html/template"
	"time"
)

type PageContent struct {
	ImagesRef models.ImagesRef
	Text      models.Text
	Texts     []models.Text
}

func humanDate(t time.Time) string {
	return t.Format("2006, jan _2 - 15h")
}

var tmpltFunctions = template.FuncMap{
	"humanDate": humanDate,
}
