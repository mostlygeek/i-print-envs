package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for _, env := range os.Environ() {
			fmt.Fprintf(w, "%s\n", env)
		}
	})

	http.HandleFunc("/__heartbeat__", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})

	var port string

	p := os.Getenv("PORT")
	if p == "" {
		log.Fatal("PORT environment variable required")
	} else {
		portNum, err := strconv.Atoi(p)

		if err != nil {
			log.Fatal(err)
		}

		if portNum < 1 || portNum > 65535 {
			log.Fatal("Invalid Port Range")
		}

		port = strconv.Itoa(portNum)
	}

	fmt.Printf("Listening on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
