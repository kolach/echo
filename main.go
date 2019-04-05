package main

import (
	"fmt"
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
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9090", nil)
}
