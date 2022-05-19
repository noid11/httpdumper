package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {

	dumpHandler := func(w http.ResponseWriter, req *http.Request) {
		body, err := httputil.DumpRequest(req, true)
		if err != nil {
			log.Fatal(err)
		}

		log.Print(string(body))
		io.WriteString(w, string(body))
	}

	http.HandleFunc("/", dumpHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
