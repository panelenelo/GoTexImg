package main

import (
	"flag"
	"net/http"
)

const port string = ":8181"

func main() {
	addr := flag.String("addr", port, "HTTP network address")
	flag.Parse()

	http.ListenAndServe(*addr, nil)
}
