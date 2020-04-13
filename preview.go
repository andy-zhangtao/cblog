package main

import (
	"fmt"
	"log"
	"net/http"
)

func startWEB(port int) {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	p := fmt.Sprintf(":%d", port)
	log.Printf("Listening on http://localhost%s ...", p)
	err := http.ListenAndServe(p, nil)
	if err != nil {
		log.Fatal(err)
	}
}
