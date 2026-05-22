package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

const port string = ":8181"

type application struct {
	PageContent PageContent
}

func main() {
	addr := flag.String("addr", port, "HTTP network address")
	flag.Parse()

	app := &application{
		PageContent: PageContent{},
	}

	log.Printf("Server starting on localhost%s", port)
	err := http.ListenAndServe(*addr, app.routes())
	if err != nil {
		log.Print(err.Error())
	}
	os.Exit(1)
}
