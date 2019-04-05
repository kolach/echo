package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	formatted, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Fprint(w, err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		fmt.Fprintf(os.Stderr, "%s\n\n", string(formatted))
		w.Write(formatted)
	}
}

func main() {
	port := flag.Int("port", 9090, "port to serve")
	flag.Parse()

	log.Printf("Starting server on port: %d\n", *port)

	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
