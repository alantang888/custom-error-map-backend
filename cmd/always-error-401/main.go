package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	listenPortEnv := os.Getenv("LISTEN_PORT")
	listenPort, err := strconv.Atoi(listenPortEnv)
	if err != nil {
		fmt.Printf("Can't parse \"%s\", use default 8080.\n", listenPortEnv)
		listenPort = 8080
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Requested header: %s\n", r.Header)

		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error 401!!"))
	})

	fmt.Printf("Server started on port %d\n", listenPort)
	http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil)
}
