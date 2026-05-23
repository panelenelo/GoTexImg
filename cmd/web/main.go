package main

import (
	"GoTexImg/internal/models"
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

const port string = ":8181"

type application struct {
	PageContent PageContent
	Client      *http.Client
	DB          models.TeximgModel
}

func main() {
	addr := flag.String("addr", port, "HTTP network address")
	flag.Parse()

	app := &application{
		PageContent: PageContent{},
		Client:      &http.Client{Timeout: 5 * time.Second},
	}

	log.Printf("Server starting on localhost%s", port)
	err := http.ListenAndServe(*addr, app.routes())
	if err != nil {
		log.Print(err.Error())
	}
	os.Exit(1)
}
